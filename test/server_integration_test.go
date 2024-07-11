package server

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"ctf01d/config"
	"ctf01d/internal/app/database"
	"ctf01d/internal/app/handlers"
	"ctf01d/internal/app/server"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/jaswdr/faker"
	_ "github.com/lib/pq"
	"github.com/tidwall/gjson"

	"github.com/go-chi/chi/v5"
)

var db *sql.DB

func TestMain(m *testing.M) {
	cfg, err := config.NewConfig("../config/config.test.yml")
	if err != nil {
		panic(err)
	}

	db, err = database.InitDatabase(cfg)
	if err != nil {
		panic(err)
	}
	code := m.Run()
	db.Close()
	os.Exit(code)
}

func NewTestRouter() (http.Handler, error) {
	hndlrs := &handlers.Handlers{
		DB: db,
	}
	svr := handlers.NewServerInterfaceWrapper(hndlrs)

	r := chi.NewRouter()
	r.Mount("/api/", server.HandlerFromMux(svr, r))
	r.Mount("/", http.HandlerFunc(server.NewHtmlRouter))

	return r, nil
}

func TestUserCRUD(t *testing.T) {
	r, err := NewTestRouter()
	if err != nil {
		t.Fatalf("failed to initialize router: %v", err)
	}

	var userID string
	// 1. Создание пользователя
	fake := faker.New()

	t.Run("Create User", func(t *testing.T) {
		user := map[string]interface{}{
			"display_name": fake.Person().Name(),
			"user_name":    fake.Gamer().Tag(),
			"role":         "player",
			"avatar_url":   "http://example.com/avatar.png",
			"status":       "active",
			"password":     fake.Internet().Password(),
		}
		body, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status code 200, got %v", rr.Code)
		}

		var response map[string]interface{}
		if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
			t.Fatalf("could not unmarshal response: %v", err)
		}

		if response["data"] != "User created successfully" {
			t.Fatalf("expected success message, got %v", response["data"])
		}
	})

	// 2. Получение всех пользователей и использование ID последнего
	t.Run("Get All Users", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/users", nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status code 200, got %v", rr.Code)
		}

		var users []map[string]interface{}
		if err := json.Unmarshal(rr.Body.Bytes(), &users); err != nil {
			t.Fatalf("could not unmarshal response: %v", err)
		}

		if len(users) == 0 {
			t.Fatalf("expected at least one user")
		}

		lastUser := users[len(users)-1]
		userID = lastUser["id"].(string)
		if userID == "" {
			t.Fatalf("expected user ID in response")
		}
	})

	// 3. Получение пользователя по ID
	t.Run("Get User by ID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/users/"+userID, nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status code 200, got %v", rr.Code)
		}

		var response map[string]interface{}
		if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
			t.Fatalf("could not unmarshal response: %v", err)
		}

		if response["id"] != userID {
			t.Fatalf("expected user ID %v, got %v", userID, response["id"])
		}
	})

	// 4. Обновление пользователя по ID
	t.Run("Update User by ID", func(t *testing.T) {
		updatedUser := map[string]interface{}{
			"display_name": fake.Person().Name(),
			"user_name":    fake.Gamer().Tag(),
			"role":         "player",
			"avatar_url":   "http://example.com/avatar_updated.png",
			"status":       "active",
			"password":     fake.Internet().Password(),
		}
		body, _ := json.Marshal(updatedUser)
		req, _ := http.NewRequest("PUT", "/api/v1/users/"+userID, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status code 200, got %v", rr.Code)
		}

		var response map[string]interface{}
		if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
			t.Fatalf("could not unmarshal response: %v", err)
		}

		if response["data"] != "User updated successfully" {
			t.Fatalf("expected success message, got %v", response["data"])
		}
	})

	// 5. Получение профиля пользователя по ID (его нет, поэтому ожидаем 404)
	t.Run("User Profile by ID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/users/"+userID+"/profile", nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Fatalf("expected status code 404, got %v", rr.Code)
		}
	})

	// 6. Удаление пользователя по ID
	t.Run("Delete User by ID", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/v1/users/"+userID, nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status code 200, got %v", rr.Code)
		}
	})
}

func TestServiceCRUD(t *testing.T) {
	r, err := NewTestRouter()
	if err != nil {
		t.Fatalf("failed to initialize router: %v", err)
	}

	var serviceID string
	faker := faker.New()

	// 1. Создание сервиса
	t.Run("Create Service", func(t *testing.T) {
		service := map[string]interface{}{
			"name":        faker.Company().Name(),
			"author":      faker.Person().Name(),
			"logo_url":    faker.Internet().URL(),
			"description": faker.Lorem().Sentence(10),
			"is_public":   faker.Bool(),
		}
		body, _ := json.Marshal(service)
		req, _ := http.NewRequest("POST", "/api/v1/services", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status code 200, got %v", rr.Code)
		}
	})

	// 2. Получение всех сервисов и использование ID последнего
	t.Run("Get All Services", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/services", nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status code 200, got %v", rr.Code)
		}

		var services []map[string]interface{}
		if err := json.Unmarshal(rr.Body.Bytes(), &services); err != nil {
			t.Fatalf("could not unmarshal response: %v", err)
		}

		if len(services) == 0 {
			t.Fatalf("expected at least one service")
		}

		lastService := services[len(services)-1]
		serviceID = lastService["id"].(string)
		if serviceID == "" {
			t.Fatalf("expected service ID in response")
		}
	})

	// 3. Получение сервиса по ID
	t.Run("Get Service by ID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/services/"+serviceID, nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status code 200, got %v", rr.Code)
		}

		var response map[string]interface{}
		if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
			t.Fatalf("could not unmarshal response: %v", err)
		}

		if response["id"] != serviceID {
			t.Fatalf("expected service ID %v, got %v", serviceID, response["id"])
		}
	})

	// 4. Обновление сервиса по ID
	t.Run("Update Service by ID", func(t *testing.T) {
		updatedService := map[string]interface{}{
			"name":        faker.Company().Name(),
			"author":      faker.Person().Name(),
			"logo_url":    faker.Internet().URL(),
			"description": faker.Lorem().Sentence(10),
			"is_public":   faker.Bool(),
		}
		body, _ := json.Marshal(updatedService)
		req, _ := http.NewRequest("PUT", "/api/v1/services/"+serviceID, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status code 200, got %v", rr.Code)
		}

		var response map[string]interface{}
		if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
			t.Fatalf("could not unmarshal response: %v", err)
		}

		if response["data"] != "Game updated successfully" {
			t.Fatalf("expected 'Game updated successfully', got %v", response["data"])
		}
	})

	// 5. Удаление сервиса по ID
	t.Run("Delete Service by ID", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/v1/services/"+serviceID, nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status code 200, got %v", rr.Code)
		}
	})
}

func TestAPIEndpoints(t *testing.T) {
	t.Skip()
	doc, err := loads.Spec("../api/openapi.yaml")
	if err != nil {
		t.Fatalf("failed to load spec: %v", err)
	}

	r, err := NewTestRouter()
	if err != nil {
		t.Fatalf("failed to initialize router: %v", err)
	}

	for path, pathItem := range doc.Spec().Paths.Paths {
		if pathItem.Get != nil {
			t.Run("GET "+path, func(t *testing.T) {
				req, err := http.NewRequest("GET", path, nil)
				if err != nil {
					t.Fatal(err)
				}

				rr := httptest.NewRecorder()

				r.ServeHTTP(rr, req)

				if status := rr.Code; status != http.StatusOK {
					t.Errorf("handler returned wrong status code for %s: got %v want %v", path, status, http.StatusOK)
				}

				expectedSchema := pathItem.Get.Responses.StatusCodeResponses[http.StatusOK].Schema

				if !validateJSONSchema(rr.Body.String(), expectedSchema) {
					t.Errorf("handler returned unexpected body for %s: got %v", path, rr.Body.String())
				}
			})
		}
	}
}

func validateJSONSchema(responseBody string, expectedSchema *spec.Schema) bool {
	expectedJSON, err := json.Marshal(expectedSchema)
	if err != nil {
		return false
	}

	return gjson.Get(responseBody, "").String() == string(expectedJSON)
}

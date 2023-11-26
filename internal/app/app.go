package app

import (
	"avitoStart/internal/endpoint"
	"avitoStart/internal/service"
	"avitoStart/internal/storage/postgres"
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

type App struct {
	database *postgres.Database
	endpoint *endpoint.Endpoint
	service  *service.Service
	echo     *echo.Echo
}

func New() (*App, error) {
	app := &App{}

	dsn := os.Getenv("DSN")

	//инициализации всех параметров через New
	app.database, _ = postgres.New(dsn)
	app.service = service.New(app.database)
	app.endpoint = endpoint.New(app.service)
	app.echo = echo.New()

	//endpoints
	app.echo.POST("/create_user", app.endpoint.AddUser)
	app.echo.POST("/delete_user", app.endpoint.DeleteUser)
	app.echo.GET("/users", app.endpoint.ExtractUsers)

	//slug
	app.echo.POST("/create_slug", app.endpoint.CreateSlug)
	app.echo.POST("/delete_slug", app.endpoint.DeleteSlug)
	//masterfunc
	app.echo.POST("/slugs_user", app.endpoint.ExecSlugNamesUser)
	app.echo.POST("/master", app.endpoint.MasterFunc)

	return app, nil
}
func (a *App) Run() error {
	log.Println("Server Runnig")

	err := a.echo.Start(":8090")
	if err != nil {
		log.Println(errors.New("Error Start Service"))
	}
	return nil
}

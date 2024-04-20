package api

import (
	"ctf01d/internal/app/models"
	"ctf01d/internal/app/repository"
	api_helpers "ctf01d/internal/app/utils"
	"ctf01d/internal/app/view"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type RequestTeam struct {
	Name         string `json:"name"`
	SocialLinks  string `json:"social_links"`
	Description  string `json:"description"`
	AvatarUrl    string `json:"avatar_url"`
	UniversityId string `json:"university_id"`
}

func CreateTeamHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var team RequestTeam
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload: " + err.Error()})
		return
	}

	teamRepo := repository.NewTeamRepository(db)
	// fixme request to model надо вынести и переиспользовать
	newTeam := &models.Team{
		Name:         team.Name,
		SocialLinks:  team.SocialLinks,
		Description:  team.Description,
		UniversityId: team.UniversityId,
		AvatarUrl:    api_helpers.PrepareImage(team.AvatarUrl),
	}
	if err := teamRepo.Create(r.Context(), newTeam); err != nil {
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create team: " + err.Error()})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Team created successfully"})
}

func DeleteTeamHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	teamRepo := repository.NewTeamRepository(db)
	if err := teamRepo.Delete(r.Context(), id); err != nil {
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete team"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Team deleted successfully"})
}

func GetTeamByIdHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	teamRepo := repository.NewTeamRepository(db)
	team, err := teamRepo.GetById(r.Context(), id)
	if err != nil {
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch team"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewTeamFromModel(team))
}

func ListTeamsHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	teamRepo := repository.NewTeamRepository(db)
	teams, err := teamRepo.List(r.Context())
	if err != nil {
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewTeamsFromModels(teams))
}

func UpdateTeamHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var team RequestTeam
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	teamRepo := repository.NewTeamRepository(db)
	updateTeam := &models.Team{
		Name:         team.Name,
		SocialLinks:  team.SocialLinks,
		Description:  team.Description,
		UniversityId: team.UniversityId,
		AvatarUrl:    team.AvatarUrl,
	}
	vars := mux.Vars(r)
	id := vars["id"]
	updateTeam.Id = id
	if err := teamRepo.Update(r.Context(), updateTeam); err != nil {
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Team updated successfully"})
}

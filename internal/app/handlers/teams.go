package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	dbmodels "ctf01d/internal/app/db"
	"ctf01d/internal/app/repository"
	"ctf01d/internal/app/server"
	api_helpers "ctf01d/internal/app/utils"
	"ctf01d/internal/app/view"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handlers) CreateTeam(w http.ResponseWriter, r *http.Request) {
	var team server.TeamRequest
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		slog.Warn(err.Error(), "handler", "CreateTeamHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}

	teamRepo := repository.NewTeamRepository(h.DB)
	newTeam := &dbmodels.Team{
		Name:         team.Name,
		SocialLinks:  *team.SocialLinks,
		Description:  *team.Description,
		UniversityId: &team.UniversityId,
		AvatarUrl:    api_helpers.PrepareImage(*team.AvatarUrl),
	}
	if err := teamRepo.Create(r.Context(), newTeam); err != nil {
		slog.Warn(err.Error(), "handler", "CreateTeamHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create team"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Team created successfully"})
}

func (h *Handlers) DeleteTeam(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	teamRepo := repository.NewTeamRepository(h.DB)
	if err := teamRepo.Delete(r.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteTeamHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete team"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Team deleted successfully"})
}

func (h *Handlers) GetTeamById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	teamRepo := repository.NewTeamRepository(h.DB)
	team, err := teamRepo.GetById(r.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetTeamByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch team"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewTeamFromModel(team))
}

func (h *Handlers) ListTeams(w http.ResponseWriter, r *http.Request) {
	teamRepo := repository.NewTeamRepository(h.DB)
	teams, err := teamRepo.List(r.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListTeamsHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to fetch teams"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewTeamsFromModels(teams))
}

func (h *Handlers) UpdateTeam(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var team server.TeamRequest
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateTeamHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	teamRepo := repository.NewTeamRepository(h.DB)
	updateTeam := &dbmodels.Team{
		Name:         team.Name,
		SocialLinks:  *team.SocialLinks,
		Description:  *team.Description,
		UniversityId: &team.UniversityId,
		AvatarUrl:    *team.AvatarUrl,
	}
	updateTeam.Id = id
	if err := teamRepo.Update(r.Context(), updateTeam); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateTeamHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to update team"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Team updated successfully"})
}

func (h *Handlers) JoinTeamUser(w http.ResponseWriter, r *http.Request, teamId openapi_types.UUID, userId openapi_types.UUID) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *Handlers) LeaveTeamUser(w http.ResponseWriter, r *http.Request, teamId openapi_types.UUID, userId openapi_types.UUID) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *Handlers) ApproveTeamUser(w http.ResponseWriter, r *http.Request, teamId openapi_types.UUID, userId openapi_types.UUID) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

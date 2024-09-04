package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"ctf01d/internal/helper"
	"ctf01d/internal/model"
	"ctf01d/internal/repository"
	"ctf01d/internal/server"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	// fixme - создание команды через апрупы - нужен стейт
	var team server.TeamRequest
	var err error
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		slog.Warn(err.Error(), "handler", "CreateTeamHandler")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}

	teamRepo := repository.NewTeamRepository(h.DB)
	newTeam := &model.Team{
		Name:         team.Name,
		SocialLinks:  helper.ToNullString(team.SocialLinks),
		Description:  *team.Description,
		UniversityId: team.UniversityId,
		AvatarUrl:    helper.ToNullString(team.AvatarUrl),
	}
	if err = teamRepo.Create(r.Context(), newTeam); err != nil {
		slog.Warn(err.Error(), "handler", "CreateTeamHandler")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create team"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, newTeam.ToResponse())
}

func (h *Handler) DeleteTeam(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	teamRepo := repository.NewTeamRepository(h.DB)
	if err := teamRepo.Delete(r.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteTeamHandler")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete team"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Team deleted successfully"})
}

func (h *Handler) GetTeamById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	teamRepo := repository.NewTeamRepository(h.DB)
	team, err := teamRepo.GetById(r.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetTeamByIdHandler")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch team"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, team.ToResponse())
}

func (h *Handler) ListTeams(w http.ResponseWriter, r *http.Request) {
	teamRepo := repository.NewTeamRepository(h.DB)
	teams, err := teamRepo.List(r.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListTeamsHandler")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to fetch teams"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, model.NewTeamsFromModels(teams))
}

func (h *Handler) UpdateTeam(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var team server.TeamRequest
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateTeamHandler")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	teamRepo := repository.NewTeamRepository(h.DB)
	updateTeam := &model.Team{
		Name:         team.Name,
		SocialLinks:  helper.ToNullString(team.SocialLinks),
		Description:  *team.Description,
		UniversityId: team.UniversityId,
		AvatarUrl:    helper.ToNullString(team.AvatarUrl),
	}
	updateTeam.Id = id
	if err := teamRepo.Update(r.Context(), updateTeam); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateTeamHandler")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to update team"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Team updated successfully"})
}

// Создает запись в таблице запросов на добавление участника в команду
func (h *Handler) ConnectUserTeam(w http.ResponseWriter, r *http.Request, teamId openapi_types.UUID, userId openapi_types.UUID) {
	// fixme move Role to openapi request
	type request struct {
		Role string `json:"role"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	teamRepo := repository.NewTeamMemberRequestRepository(h.DB)

	err := teamRepo.ConnectUserTeam(r.Context(), teamId, userId, req.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User create request to join the team"})
}

// Обновляет запись в таблице запросов и добавляет пользователя в команду
func (h *Handler) ApproveUserTeam(w http.ResponseWriter, r *http.Request, teamId openapi_types.UUID, userId openapi_types.UUID) {
	teamRepo := repository.NewTeamMemberRequestRepository(h.DB)
	err := teamRepo.ApproveUserTeam(r.Context(), teamId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User approved and added to team successfully"})
}

// Удаляет пользователя из команды
func (h *Handler) LeaveUserFromTeam(w http.ResponseWriter, r *http.Request, teamId openapi_types.UUID, userId openapi_types.UUID) {
	teamRepo := repository.NewTeamMemberRequestRepository(h.DB)
	err := teamRepo.LeaveUserFromTeam(r.Context(), teamId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User removed from team successfully"})
}

func (h *Handler) TeamMembers(w http.ResponseWriter, r *http.Request, teamId openapi_types.UUID) {
	// нужно по парамтрам иметь возможноть получать pending user'ов для опрува
	teamRepo := repository.NewTeamMemberRequestRepository(h.DB)
	members, err := teamRepo.TeamMembers(r.Context(), teamId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, model.NewUsersFromModels(members))
}

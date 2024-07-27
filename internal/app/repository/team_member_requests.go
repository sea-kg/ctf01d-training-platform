package repository

import (
	"context"
	models "ctf01d/internal/app/db"
	"database/sql"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type TeamMemberRequestRepository interface {
	ConnectUserTeam(ctx context.Context, teamID, userID openapi_types.UUID, role string) error
	ApproveUserTeam(ctx context.Context, teamID, userID openapi_types.UUID) error
	LeaveUserFromTeam(ctx context.Context, teamID, userID openapi_types.UUID) error
	TeamMembers(ctx context.Context, teamID openapi_types.UUID) ([]*models.User, error)
}

func NewTeamMemberRequestRepository(db *sql.DB) TeamMemberRequestRepository {
	return &teamRepo{db: db}
}

func (r *teamRepo) ConnectUserTeam(ctx context.Context, teamID, userID openapi_types.UUID, role string) error {
	query := `INSERT INTO team_member_requests (team_id, user_id, role, status)
	          VALUES ($1, $2, $3, 'pending')`
	_, err := r.db.ExecContext(ctx, query, teamID, userID, role)
	return err
}

func (r *teamRepo) ApproveUserTeam(ctx context.Context, teamID, userID openapi_types.UUID) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := `UPDATE team_member_requests SET status = 'approved' WHERE team_id = $1 AND user_id = $2 AND status = 'pending'`
	_, err = tx.ExecContext(ctx, query, teamID, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	var role string
	query = `SELECT role FROM team_member_requests WHERE team_id = $1 AND user_id = $2 AND status = 'approved'`
	err = tx.QueryRowContext(ctx, query, teamID, userID).Scan(&role)
	if err != nil {
		tx.Rollback()
		return err
	}
	// fixme - обновить team_history
	query = `INSERT INTO profiles (current_team_id, user_id, role) VALUES ($1, $2, $3)`
	_, err = tx.ExecContext(ctx, query, teamID, userID, role)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *teamRepo) LeaveUserFromTeam(ctx context.Context, teamID, userID openapi_types.UUID) error {
	// fixme - обновить team_history
	query := `DELETE FROM profiles WHERE current_team_id = $1 AND user_id = $2`
	_, err := r.db.ExecContext(ctx, query, teamID, userID)
	return err
}

func (r *teamRepo) TeamMembers(ctx context.Context, teamID openapi_types.UUID) ([]*models.User, error) {
	query := `SELECT u.id, u.display_name, u.user_name, tm.role, u.avatar_url, u.status
	          FROM profiles tm
	          JOIN users u ON tm.user_id = u.id
	          WHERE tm.current_team_id = $1`
	rows, err := r.db.QueryContext(ctx, query, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []*models.User
	for rows.Next() {
		var member models.User
		if err := rows.Scan(&member.Id, &member.DisplayName, &member.Username, &member.Role, &member.AvatarUrl, &member.Status); err != nil {
			return nil, err
		}
		members = append(members, &member)
	}
	return members, nil
}

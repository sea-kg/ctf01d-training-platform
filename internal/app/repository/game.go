package repository

import (
	"context"
	models "ctf01d/internal/app/db"
	"database/sql"
	"time"
)

type GameRepository interface {
	Create(ctx context.Context, game *models.Game) error
	GetById(ctx context.Context, id int) (*models.Game, error)
	GetGameDetails(ctx context.Context, id int) (*models.GameDetails, error)
	Update(ctx context.Context, game *models.Game) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*models.Game, error)
}

type gameRepo struct {
	db *sql.DB
}

func NewGameRepository(db *sql.DB) GameRepository {
	return &gameRepo{db: db}
}

func (r *gameRepo) Create(ctx context.Context, game *models.Game) error {
	query := `INSERT INTO games (start_time, end_time, description) VALUES ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query, game.StartTime, game.EndTime, game.Description)
	return err
}

func (r *gameRepo) GetGameDetails(ctx context.Context, id int) (*models.GameDetails, error) {
	query := `
        SELECT g.id, g.start_time, g.end_time, g.description, t.id, t.name, t.description, u.id, u.user_name
        FROM games g
        JOIN team_games tg ON g.id = tg.game_id
        JOIN teams t ON tg.team_id = t.id
        JOIN team_members tm ON t.id = tm.team_id
        JOIN users u ON tm.user_id = u.id
        WHERE g.id = $1;
    `
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	gameDetails := &models.GameDetails{}
	teams := map[int]*models.TeamDetails{}

	for rows.Next() {
		var gameId int
		var startTime, endTime time.Time
		var description string
		var teamId int
		var teamName string
		var teamDescription string
		var userId int
		var userName string

		err := rows.Scan(&gameId, &startTime, &endTime, &description, &teamId, &teamName, &teamDescription, &userId, &userName)
		if err != nil {
			return nil, err
		}

		if gameDetails.Id == 0 {
			gameDetails.Id = gameId
			gameDetails.StartTime = startTime
			gameDetails.EndTime = endTime
			gameDetails.Description = description
		}

		if team, ok := teams[teamId]; ok {
			team.Members = append(team.Members, &models.User{Id: userId, Username: userName})
		} else {
			newTeam := &models.TeamDetails{
				Team: models.Team{
					Id:          teamId,
					Name:        teamName,
					Description: teamDescription,
				},
				Members: []*models.User{{Id: userId, Username: userName}},
			}

			teams[teamId] = newTeam
			gameDetails.TeamDetails = append(gameDetails.TeamDetails, newTeam)
		}
	}
	return gameDetails, nil
}

func (r *gameRepo) GetById(ctx context.Context, id int) (*models.Game, error) {
	query := `SELECT id, start_time, end_time, description FROM games WHERE id = $1`
	game := &models.Game{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&game.Id, &game.StartTime, &game.EndTime, &game.Description)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (r *gameRepo) Update(ctx context.Context, game *models.Game) error {
	query := `UPDATE games SET start_time = $1, end_time = $2, description = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, game.StartTime, game.EndTime, game.Description, game.Id)
	return err
}

func (r *gameRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM games WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *gameRepo) List(ctx context.Context) ([]*models.Game, error) {
	query := `SELECT id, start_time, end_time, description FROM games`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []*models.Game
	for rows.Next() {
		var game models.Game
		if err := rows.Scan(&game.Id, &game.StartTime, &game.EndTime, &game.Description); err != nil {
			return nil, err
		}
		games = append(games, &game)
	}
	return games, nil
}

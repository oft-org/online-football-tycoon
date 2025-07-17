package tournament

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_tournament_by_season_id.sql
var getTournamentBySeasonIDQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	getTournamentBySeasonIDStmt, err := db.Prepare(getTournamentBySeasonIDQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                      db,
		getTournamentBySeasonID: getTournamentBySeasonIDStmt,
	}, nil
}

type Repository struct {
	db                      *sql.DB
	getTournamentBySeasonID *sql.Stmt
}

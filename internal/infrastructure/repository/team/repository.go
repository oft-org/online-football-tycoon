package team

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_season_team.sql
var getSeasonTeamQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	getSeasonTeamStmt, err := db.Prepare(getSeasonTeamQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:            db,
		getSeasonTeam: getSeasonTeamStmt,
	}, nil
}

type Repository struct {
	db            *sql.DB
	getSeasonTeam *sql.Stmt
}

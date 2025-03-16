package repository

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_matches.sql
var getMatchesQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	getMatchesStmt, err := db.Prepare(getMatchesQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:         db,
		getMatches: getMatchesStmt,
	}, nil
}

type repository struct {
	db         *sql.DB
	getMatches *sql.Stmt
}

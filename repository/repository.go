package repository

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_matches.sql
var getMatchesQuery string

//go:embed sql/get_matches.sql
var getMatchByIdQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	getMatchesStmt, err := db.Prepare(getMatchesQuery)
	if err != nil {
		return nil, err
	}

	getMatchByIdStmt, err := db.Prepare(getMatchByIdQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:           db,
		getMatches:   getMatchesStmt,
		getMatchById: getMatchByIdStmt,
	}, nil
}

type repository struct {
	db           *sql.DB
	getMatches   *sql.Stmt
	getMatchById *sql.Stmt
}

package repository

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_matches.sql
var getMatchesQuery string

//go:embed sql/get_match_by_id.sql
var getMatchByIdQuery string

//go:embed sql/post_match.sql
var postMatchQuery string

//go:embed sql/post_match_events.sql
var postMatchEventsQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	getMatchesStmt, err := db.Prepare(getMatchesQuery)
	if err != nil {
		return nil, err
	}

	getMatchByIdStmt, err := db.Prepare(getMatchByIdQuery)
	if err != nil {
		return nil, err
	}

	postMatchStmt, err := db.Prepare(postMatchQuery)
	if err != nil {
		return nil, err
	}

	postMatchEventsStmt, err := db.Prepare(postMatchEventsQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:              db,
		getMatches:      getMatchesStmt,
		getMatchById:    getMatchByIdStmt,
		postMatch:       postMatchStmt,
		postMatchEvents: postMatchEventsStmt,
	}, nil
}

type repository struct {
	db              *sql.DB
	getMatches      *sql.Stmt
	getMatchById    *sql.Stmt
	postMatch       *sql.Stmt
	postMatchEvents *sql.Stmt
}

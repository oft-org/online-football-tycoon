package match

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/post_player.sql
var postPlayerQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	postPlayerStmt, err := db.Prepare(postPlayerQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:         db,
		postPlayer: postPlayerStmt,
	}, nil
}

type Repository struct {
	db         *sql.DB
	postPlayer *sql.Stmt
}

package classification

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_classification.sql
var getClassificationQuery string

func NewRepository(db *sql.DB) (*Repository, error) {

	getClassificationStmt, err := db.Prepare(getClassificationQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                db,
		getClassification: getClassificationStmt,
	}, nil
}

type Repository struct {
	db                *sql.DB
	getClassification *sql.Stmt
}

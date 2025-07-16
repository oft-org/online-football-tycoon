package classification

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_classification.sql
var getClassificationQuery string

//go:embed sql/get_team_classification.sql
var getTeamClassificationQuery string

//go:embed sql/insert_classification.sql
var insertClassificationQuery string

//go:embed sql/update_classification.sql
var updateClassificationQuery string

func NewRepository(db *sql.DB) (*Repository, error) {

	getClassificationStmt, err := db.Prepare(getClassificationQuery)
	if err != nil {
		return nil, err
	}

	getTeamClassificationStmt, err := db.Prepare(getTeamClassificationQuery)
	if err != nil {
		return nil, err
	}

	insertClassificationStmt, err := db.Prepare(insertClassificationQuery)
	if err != nil {
		return nil, err
	}

	updateClassificationStmt, err := db.Prepare(updateClassificationQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                    db,
		getClassification:     getClassificationStmt,
		getTeamClassification: getTeamClassificationStmt,
		insertClassification:  insertClassificationStmt,
		updateClassification:  updateClassificationStmt,
	}, nil
}

type Repository struct {
	db                    *sql.DB
	getClassification     *sql.Stmt
	getTeamClassification *sql.Stmt
	insertClassification  *sql.Stmt
	updateClassification  *sql.Stmt
}

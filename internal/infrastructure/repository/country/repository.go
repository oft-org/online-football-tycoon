package country

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_countries.sql
var getCountriesQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	getCountriesStmt, err := db.Prepare(getCountriesQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:           db,
		getCountries: getCountriesStmt,
	}, nil
}

type Repository struct {
	db           *sql.DB
	getCountries *sql.Stmt
}

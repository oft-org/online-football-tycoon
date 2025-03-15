package repository

import "database/sql"

func NewRepository(db *sql.DB) (*repository, error) {
	return &repository{
		db: db,
	}, nil
}

type repository struct {
	db *sql.DB
}

package match

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestGetMatches(t *testing.T) {

	connStr := "user=postgres password=mysecretpassword dbname=online_football_tycoon host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	repo, err := NewRepository(db)
	if err != nil {
		t.Fatalf("error al crear el repositorio: %v", err)
	}

	matches, err := repo.GetMatches()

	assert.NoError(t, err)

	assert.NotEmpty(t, matches)
}

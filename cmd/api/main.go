package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/robertobouses/online-football-tycoon/http"
	"github.com/robertobouses/online-football-tycoon/match"
	"github.com/robertobouses/online-football-tycoon/repository"
)

type DBConfig struct {
	User     string
	Pass     string
	Host     string
	Port     string
	Database string
}

func NewPostgres(c DBConfig) (*sql.DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Pass, c.Host, c.Port, c.Database)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Println("error en apertura de conexión")
		return nil, err
	}
	log.Println("la conexión tiene las credenciales correctas")
	return db, db.Ping()
}

func main() {
	db, err := NewPostgres(DBConfig{
		User:     "postgres",
		Pass:     "mysecretpassword",
		Host:     "localhost",
		Port:     "5432",
		Database: "online_football_tycoon",
	})

	if err != nil {
		log.Println(err)
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM oft.match")
	log.Println("**/*/***/*/*/*/*rows", rows)

	repository, err := repository.NewRepository(db)
	if err != nil {
		panic(err)
	}
	app := match.NewApp(repository)

	handler := http.NewHandler(app)

	s := http.NewServer(handler)

	s.Run("8080")
}

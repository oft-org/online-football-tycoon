package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/robertobouses/online-football-tycoon/http"
	"github.com/robertobouses/online-football-tycoon/internal"
	"github.com/robertobouses/online-football-tycoon/match"
	"github.com/robertobouses/online-football-tycoon/repository"
)

func main() {

	requiredEnv := []string{"DB_USER", "DB_PASS", "DB_HOST", "DB_PORT", "DB_NAME"}
	for _, env := range requiredEnv {
		if os.Getenv(env) == "" {
			log.Fatalf("missing required environment variable: %s", env)
		}
	}
	db, err := internal.NewPostgres(internal.DBConfig{
		User:     os.Getenv("DB_USER"),
		Pass:     os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
	})

	if err != nil {
		log.Println(err)
		panic(err)
	}

	repository, err := repository.NewRepository(db)
	if err != nil {
		panic(err)
	}
	app := match.NewApp(repository)

	handler := http.NewHandler(app)

	s := http.NewServer(handler)

	if err := s.Run("8080"); err != nil {
		log.Printf("error running server: %v\n", err)
	}
}

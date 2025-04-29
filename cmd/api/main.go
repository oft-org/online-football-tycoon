package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/robertobouses/online-football-tycoon/http"
	"github.com/robertobouses/online-football-tycoon/internal"
	"github.com/robertobouses/online-football-tycoon/match"
	"github.com/robertobouses/online-football-tycoon/repository"
)

func main() {
	db, err := internal.NewPostgres(internal.DBConfig{
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

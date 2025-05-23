package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/robertobouses/online-football-tycoon/http"
	"github.com/robertobouses/online-football-tycoon/internal"
	"github.com/robertobouses/online-football-tycoon/match"
	"github.com/robertobouses/online-football-tycoon/repository"
)

func main() {
	err := godotenv.Load(".env.config")
	if err != nil {
		log.Fatalf("error loading .env.config: %v", err)
	}

	fmt.Println("DB_USER:", os.Getenv("DB_USER"))

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
	fmt.Println("Step 1: Connection to the database created")

	if err != nil {
		log.Println(err)
		panic(err)
	}

	repository, err := repository.NewRepository(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Step 2: repository created")

	app := match.NewApp(repository)
	fmt.Println("Step 3: app created")

	handler := http.NewHandler(app)
	fmt.Println("Step 4: Handler created")

	s := http.NewServer(handler)
	fmt.Println("Step 5: Server created")

	log.Println("Server starting on port 8080...")
	if err := s.Run("8080"); err != nil {
		log.Printf("error running server: %v\n", err)
	}
}

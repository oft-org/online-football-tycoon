package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/robertobouses/online-football-tycoon/http"
	"github.com/robertobouses/online-football-tycoon/internal"
	"github.com/robertobouses/online-football-tycoon/match"
	"github.com/robertobouses/online-football-tycoon/repository"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the API server",
	Run: func(cmd *cobra.Command, args []string) {
		_ = godotenv.Load(".env.config")

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
			log.Fatal("failed to connect to database:", err)
		}

		repo, err := repository.NewRepository(db)
		if err != nil {
			log.Fatal("failed to init repository:", err)
		}

		app := match.NewApp(repo)
		handler := http.NewHandler(app)
		s := http.NewServer(handler)

		if err := s.Run("8080"); err != nil {
			log.Fatal("server failed:", err)
		}
	},
}

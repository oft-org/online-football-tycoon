package server

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/robertobouses/online-football-tycoon/internal/domain/use_cases/match"
	"github.com/robertobouses/online-football-tycoon/internal/infrastructure/http"
	"github.com/robertobouses/online-football-tycoon/internal/infrastructure/http/handler"
	repository "github.com/robertobouses/online-football-tycoon/internal/infrastructure/repository/match"

	internalPostgres "github.com/robertobouses/online-football-tycoon/internal/pkg/postgres"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the API server",
	Run: func(cmd *cobra.Command, args []string) {
		_ = godotenv.Load()

		requiredEnv := []string{"DB_USER", "DB_PASS", "DB_HOST", "DB_PORT", "DB_NAME"}
		for _, env := range requiredEnv {
			if os.Getenv(env) == "" {
				log.Fatalf("missing required environment variable: %s", env)
			}
		}

		db, err := internalPostgres.NewPostgres(internalPostgres.DBConfig{
			User:     os.Getenv("DB_USER"),
			Pass:     os.Getenv("DB_PASS"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_NAME"),
		})
		if err != nil {
			log.Fatal("failed to connect to database:", err)
		}
		// TODO: ahora aqui, este repository.NewRepository tiene que ser el repository que esta en internal/infraestructure/repositopry/match/reposiutory.go
		// DONE: Ya lo es?? No comprendo
		repo, err := repository.NewRepository(db)
		if err != nil {
			log.Fatal("failed to init repository:", err)
		}

		app := match.NewApp(repo)
		handler := handler.NewHandler(app)
		s := http.NewServer(handler)

		if err := s.Run("8080"); err != nil {
			log.Fatal("server failed:", err)
		}
	},
}

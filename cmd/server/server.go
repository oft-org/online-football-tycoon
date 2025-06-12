package server

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	appMatch "github.com/robertobouses/online-football-tycoon/internal/domain/use_cases/match"
	appPlayer "github.com/robertobouses/online-football-tycoon/internal/domain/use_cases/player"
	appTeam "github.com/robertobouses/online-football-tycoon/internal/domain/use_cases/team"
	httpServer "github.com/robertobouses/online-football-tycoon/internal/infrastructure/http"
	handlerMatch "github.com/robertobouses/online-football-tycoon/internal/infrastructure/http/match"
	handlerPlayer "github.com/robertobouses/online-football-tycoon/internal/infrastructure/http/player"
	repositoryMatch "github.com/robertobouses/online-football-tycoon/internal/infrastructure/repository/match"
	repositoryPlayer "github.com/robertobouses/online-football-tycoon/internal/infrastructure/repository/player"
	repositoryTeam "github.com/robertobouses/online-football-tycoon/internal/infrastructure/repository/team"
	internalPostgres "github.com/robertobouses/online-football-tycoon/internal/pkg/postgres"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the API server",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("failed to get env:", err)
		}

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
		matchRepo, err := repositoryMatch.NewRepository(db)
		if err != nil {
			log.Fatal("failed to init repository:", err)
		}
		playerRepo, err := repositoryPlayer.NewRepository(db)
		if err != nil {
			log.Fatal("failed to init repository:", err)
		}
		teamRepo, err := repositoryTeam.NewRepository(db)
		if err != nil {
			log.Fatal("failed to init repository:", err)
		}
		matchApp := appMatch.NewApp(matchRepo)
		playerApp := appPlayer.NewApp(playerRepo)
		teamApp := appTeam.NewApp(teamRepo, *matchRepo)

		matchHandler := handlerMatch.NewHandler(matchApp, teamApp)
		playerHandler := handlerPlayer.NewHandler(playerApp)
		s := httpServer.NewServer(matchHandler, playerHandler)

		if err := s.Run("8080"); err != nil {
			log.Fatal("server failed:", err)
		}
	},
}

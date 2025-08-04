package server

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	appClassification "github.com/robertobouses/online-football-tycoon/internal/domain/use_cases/classification"
	appCountry "github.com/robertobouses/online-football-tycoon/internal/domain/use_cases/country"
	appMatch "github.com/robertobouses/online-football-tycoon/internal/domain/use_cases/match"
	appPlayer "github.com/robertobouses/online-football-tycoon/internal/domain/use_cases/player"
	appTournament "github.com/robertobouses/online-football-tycoon/internal/domain/use_cases/tournament"
	httpServer "github.com/robertobouses/online-football-tycoon/internal/infrastructure/http"
	handlerClassification "github.com/robertobouses/online-football-tycoon/internal/infrastructure/http/classification"
	handlerCountry "github.com/robertobouses/online-football-tycoon/internal/infrastructure/http/country"
	handlerMatch "github.com/robertobouses/online-football-tycoon/internal/infrastructure/http/match"
	handlerPlayer "github.com/robertobouses/online-football-tycoon/internal/infrastructure/http/player"
	handlerTournament "github.com/robertobouses/online-football-tycoon/internal/infrastructure/http/tournament"
	repositoryClassification "github.com/robertobouses/online-football-tycoon/internal/infrastructure/repository/classification"
	repositoryCountry "github.com/robertobouses/online-football-tycoon/internal/infrastructure/repository/country"
	repositoryMatch "github.com/robertobouses/online-football-tycoon/internal/infrastructure/repository/match"
	repositoryPlayer "github.com/robertobouses/online-football-tycoon/internal/infrastructure/repository/player"
	repositoryTeam "github.com/robertobouses/online-football-tycoon/internal/infrastructure/repository/team"
	repositoryTournament "github.com/robertobouses/online-football-tycoon/internal/infrastructure/repository/tournament"
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
			log.Fatal("failed to init match repository:", err)
		}
		playerRepo, err := repositoryPlayer.NewRepository(db)
		if err != nil {
			log.Fatal("failed to init player repository:", err)
		}
		teamRepo, err := repositoryTeam.NewRepository(db)
		if err != nil {
			log.Fatal("failed to init team repository:", err)
		}
		classificationRepo, err := repositoryClassification.NewRepository(db)
		if err != nil {
			log.Fatal("failde to init classification repository:", err)
		}
		tournamentRepo, err := repositoryTournament.NewRepository(db)
		if err != nil {
			log.Fatal("failde to init tournament repository:", err)
		}
		countryRepo, err := repositoryCountry.NewRepository(db)
		if err != nil {
			log.Fatal("failde to init country repository:", err)

		}

		matchApp := appMatch.NewApp(matchRepo, classificationRepo, teamRepo)
		playerApp := appPlayer.NewApp(playerRepo)
		classificationApp := appClassification.NewApp(classificationRepo, tournamentRepo)
		countryApp := appCountry.NewApp(countryRepo)
		tournamentApp := appTournament.NewApp(tournamentRepo, teamRepo, matchRepo)

		matchHandler := handlerMatch.NewHandler(matchApp, &tournamentApp)
		playerHandler := handlerPlayer.NewHandler(playerApp)
		classificationHandler := handlerClassification.NewHandler(classificationApp)
		countryHandler := handlerCountry.NewHandler(countryApp)
		tournamentHandler := handlerTournament.NewHandler(tournamentApp)

		s := httpServer.NewServer(matchHandler, playerHandler, classificationHandler, countryHandler, *tournamentHandler)

		if err := s.Run("8080"); err != nil {
			log.Fatal("server failed:", err)
		}
	},
}

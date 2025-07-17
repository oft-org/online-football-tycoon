package team

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/infrastructure/repository/match"
	"github.com/robertobouses/online-football-tycoon/internal/infrastructure/repository/tournament"
)

type Repository interface {
	GetSeasonTeam(seasonID uuid.UUID) ([]uuid.UUID, error)
}

func NewApp(repository Repository, matchRepo match.Repository, tournamentRepo tournament.Repository) AppService {
	return AppService{
		repo:           repository,
		matchRepo:      matchRepo,
		tournamentRepo: tournamentRepo,
	}
}

type AppService struct {
	repo           Repository
	matchRepo      match.Repository
	tournamentRepo tournament.Repository
}

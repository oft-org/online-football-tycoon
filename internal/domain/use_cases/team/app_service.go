package team

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/infrastructure/repository/match"
)

type Repository interface {
	GetSeasonTeam(seasonID uuid.UUID) ([]uuid.UUID, error)
}

func NewApp(repository Repository, matchRepo match.Repository) AppService {
	return AppService{
		repo:      repository,
		matchRepo: matchRepo,
	}
}

type AppService struct {
	repo      Repository
	matchRepo match.Repository
}

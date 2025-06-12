package team

import (
	"github.com/robertobouses/online-football-tycoon/internal/domain"
	"github.com/robertobouses/online-football-tycoon/internal/infrastructure/repository/match"
)

type Repository interface {
	GetSeasonTeam() ([]domain.SeasonTeam, error)
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

package tournament

import (
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type Repository interface {
	GetTournamentsByCountry(country string) ([]domain.Tournament, error)
}

func NewApp(repository Repository) AppService {
	return AppService{
		repo: repository,
	}
}

type AppService struct {
	repo Repository
}

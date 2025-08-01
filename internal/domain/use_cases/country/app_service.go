package country

import (
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type Repository interface {
	GetCountries() ([]domain.Country, error)
}

func NewApp(repository Repository) AppService {
	return AppService{
		repo: repository,
	}
}

type AppService struct {
	repo Repository
}

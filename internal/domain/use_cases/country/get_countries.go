package country

import (
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (a AppService) GetCountries() ([]domain.Country, error) {
	countries, err := a.repo.GetCountries()
	if err != nil {
		return []domain.Country{}, err
	}
	return countries, nil

}

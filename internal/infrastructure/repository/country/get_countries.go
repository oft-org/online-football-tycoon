package country

import (
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) GetCountries() ([]domain.Country, error) {
	rows, err := r.getCountries.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var countries []domain.Country
	for rows.Next() {
		var country domain.Country
		err := rows.Scan(
			&country.Code,
			&country.Continent,
		)
		if err != nil {
			return nil, err
		}

		countries = append(countries, country)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return countries, nil
}

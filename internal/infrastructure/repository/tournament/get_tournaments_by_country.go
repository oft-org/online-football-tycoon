package tournament

import (
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) GetTournamentsByCountry(country string) ([]domain.Tournament, error) {
	rows, err := r.getTournamentsByCountry.Query(country)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tournaments []domain.Tournament

	for rows.Next() {
		var tournament domain.Tournament
		if err := rows.Scan(
			&tournament.ID,
			&tournament.Name,
			&tournament.Type,
			&tournament.CountryCode,
			&tournament.Division,
			&tournament.PromotionTo,
			&tournament.DescentTo,
		); err != nil {
			return nil, err
		}
		tournaments = append(tournaments, tournament)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tournaments, nil
}

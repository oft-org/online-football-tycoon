package team

import (
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) GetSeasonTeam() ([]domain.SeasonTeam, error) {
	rows, err := r.getSeasonTeam.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seasonTeams []domain.SeasonTeam
	for rows.Next() {
		var seasonTeam domain.SeasonTeam
		err := rows.Scan(
			&seasonTeam.SeasonID,
			&seasonTeam.TeamID,
		)
		if err != nil {
			return nil, err
		}
		seasonTeams = append(seasonTeams, seasonTeam)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return seasonTeams, nil
}

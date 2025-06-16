package team

import (
	"github.com/google/uuid"
)

func (r *Repository) GetSeasonTeam(seasonId uuid.UUID) ([]uuid.UUID, error) {
	rows, err := r.getSeasonTeam.Query(seasonId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seasonTeams []uuid.UUID

	for rows.Next() {
		var teamID uuid.UUID
		err := rows.Scan(&teamID)
		if err != nil {
			return nil, err
		}
		seasonTeams = append(seasonTeams, teamID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return seasonTeams, nil
}

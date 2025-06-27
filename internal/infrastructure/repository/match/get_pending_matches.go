package match

import (
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) GetPendingMatches(timestamp time.Time) ([]domain.SeasonMatch, error) {
	rows, err := r.getPendingMatches.Query(timestamp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []domain.SeasonMatch
	for rows.Next() {
		var m domain.SeasonMatch
		var matchID, seasonID, homeTeam, awayTeam uuid.UUID
		var matchDate time.Time
		var homeResult, awayResult int
		err := rows.Scan(
			&matchID,
			&seasonID,
			&homeTeam,
			&awayTeam,
			&matchDate,
			&homeResult,
			&awayResult,
		)
		if err != nil {
			return nil, err
		}
		m.ID = matchID
		m.SeasonID = seasonID
		m.HomeTeamID = homeTeam
		m.AwayTeamID = awayTeam
		m.MatchDate = matchDate
		m.HomeResult = &homeResult
		m.AwayResult = &awayResult

		matches = append(matches, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

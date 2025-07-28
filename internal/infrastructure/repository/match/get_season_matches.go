package match

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r Repository) GetSeasonMatches(seasonID uuid.UUID) ([]domain.SeasonMatch, error) {
	rows, err := r.getSeasonMatches.Query(seasonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []domain.SeasonMatch
	for rows.Next() {
		var m domain.SeasonMatch
		var matchID, seasonID, homeTeam, awayTeam uuid.UUID
		var matchDate time.Time
		var homeResult, awayResult sql.NullInt32

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

		if homeResult.Valid {
			val := int(homeResult.Int32)
			m.HomeResult = &val
		} else {
			m.HomeResult = nil
		}

		if awayResult.Valid {
			val := int(awayResult.Int32)
			m.AwayResult = &val
		} else {
			m.AwayResult = nil
		}

		matches = append(matches, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

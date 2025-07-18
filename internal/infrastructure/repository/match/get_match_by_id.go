package match

import (
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) GetMatchByID(matchID uuid.UUID) (*domain.SeasonMatch, error) {
	row := r.getMatchByID.QueryRow(matchID)

	var match domain.SeasonMatch
	err := row.Scan(
		&match.ID,
		&match.SeasonID,
		&match.HomeTeamID,
		&match.AwayTeamID,
		&match.MatchDate,
		&match.HomeResult,
		&match.AwayResult,
	)

	log.Printf("GetMatchByID returned match: ID=%v, HomeResult=%v, AwayResult=%v", match.ID, match.HomeResult, match.AwayResult)

	if err != nil {
		return nil, err
	}

	return &match, nil
}

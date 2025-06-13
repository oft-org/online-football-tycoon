package match

import (
	"log"
	"time"

	"github.com/google/uuid"
)

func (r *Repository) PostMatch(seasonId, homeTeamId, awayTeamId uuid.UUID, matchDate time.Time, homeGoals, awayGoals int) error {

	_, err := r.postMatch.Exec(
		seasonId,
		homeTeamId,
		awayTeamId,
		matchDate,
		homeGoals,
		awayGoals,
	)

	if err != nil {
		log.Print("Error executing PostMatch statement:", err)
		return err
	}

	return nil
}

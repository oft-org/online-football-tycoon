package match

import (
	"errors"
	"log"

	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) UpdateMatch(seasonMatch domain.SeasonMatch) error {
	match, err := r.GetMatchByID(seasonMatch.ID)
	if err != nil {
		return err
	}

	if match.HomeResult != nil || match.AwayResult != nil {
		return errors.New("match already played, cannot update")
	}

	_, err = r.updateMatch.Exec(
		seasonMatch.ID,
		seasonMatch.HomeResult,
		seasonMatch.AwayResult,
	)
	if err != nil {
		log.Print("Error executing UpdateMatch statement:", err)
		return err
	}

	return nil
}

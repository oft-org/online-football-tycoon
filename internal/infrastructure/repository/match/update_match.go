package match

import (
	"errors"
	"log"

	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) UpdateMatch(seasonMatch domain.SeasonMatch) error {
	match, err := r.GetMatchByID(seasonMatch.ID)
	if err != nil {
		log.Println("Error GetMatchByID in UpdateMatch")
		return err
	}

	log.Printf("Checking if match already played: HomeResult=%v, AwayResult=%v", match.HomeResult, match.AwayResult)
	if match.HomeResult != nil || match.AwayResult != nil {
		return errors.New("match already played, cannot update")
	}

	if r.updateMatch == nil {
		log.Fatal("updateMatch is nil — SQL statement was not prepared")
	}
	_, err = r.updateMatch.Exec(
		seasonMatch.ID,
		seasonMatch.HomeResult,
		seasonMatch.AwayResult,
	)
	log.Println("UpdateMatch after Exec")
	if err != nil {
		log.Print("Error executing UpdateMatch statement:", err)
		return err
	}

	return nil
}

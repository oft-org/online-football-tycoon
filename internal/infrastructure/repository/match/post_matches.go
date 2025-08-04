package match

import (
	"log"

	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) PostMatches(matches []domain.SeasonMatch) error {
	for _, match := range matches {

		if r.postMatch == nil {
			log.Fatal("panic prevented: r.postMatch is nil in Repository")
		}
		_, err := r.postMatch.Exec(
			match.SeasonID,
			match.HomeTeamID,
			match.AwayTeamID,
			match.MatchDate,
			match.HomeResult,
			match.AwayResult,
		)
		if err != nil {
			log.Printf("Error inserting match %v: %v", match.ID, err)
			return err
		}
	}

	return nil
}

package classification

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) UpdateClassification(classification domain.Classification) error {
	team, err := r.GetTeamClassification(classification.TeamID)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Classification not found for team %s, inserting new record", classification.TeamID)
			return r.InsertClassification(classification)
		}
		return fmt.Errorf("error fetching current team classification: %w", err)
	}

	if classification.Points <= 0 {
		classification.Points = team.Points
	}
	if classification.GoalsFor <= 0 {
		classification.GoalsFor = team.GoalsFor
	}
	if classification.GoalsAgainst <= 0 {
		classification.GoalsAgainst = team.GoalsAgainst
	}

	_, err = r.updateClassification.Exec(
		classification.TeamID,
		classification.Points,
		classification.GoalsFor,
		classification.GoalsAgainst)
	if err != nil {
		log.Printf("Error updating classification: %v", err)
		return err
	}
	return nil
}

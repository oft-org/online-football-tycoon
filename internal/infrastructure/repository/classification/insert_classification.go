package classification

import (
	"log"

	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) InsertClassification(classification domain.Classification) error {
	_, err := r.insertClassification.Exec(classification.TeamID, classification.Points, classification.GoalsFor, classification.GoalsAgainst)
	if err != nil {
		log.Printf("Error inserting classification: %v", err)
	}
	return err
}

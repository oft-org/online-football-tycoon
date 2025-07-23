package classification

import (
	"log"

	"github.com/google/uuid"
)

type ClassificationWithTournament struct {
	TeamID         uuid.UUID
	TeamName       string
	Position       int
	Points         int
	GoalsFor       int
	GoalsAgainst   int
	GoalDifference int
	TournamentName string
	Country        string
}

func (a AppService) GetClassification(seasonID uuid.UUID) ([]ClassificationWithTournament, error) {
	classification, err := a.classificationRepo.GetClassification(seasonID)
	if err != nil {
		return nil, err
	}

	tournament, err := a.tournamentRepo.GetTournamentBySeasonID(seasonID)
	if err != nil {
		log.Println("Error Get Tournament on GetClassification")
		return nil, err
	}

	var result []ClassificationWithTournament
	for _, c := range classification {
		result = append(result, ClassificationWithTournament{
			TeamID:         c.TeamID,
			TeamName:       c.TeamName,
			Position:       c.Position,
			Points:         c.Points,
			GoalsFor:       c.GoalsFor,
			GoalsAgainst:   c.GoalsAgainst,
			GoalDifference: c.GoalDifference,
			TournamentName: tournament.Name,
			Country:        tournament.CountryCode,
		})
	}

	return result, nil
}

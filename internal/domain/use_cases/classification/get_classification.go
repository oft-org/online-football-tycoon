package classification

import (
	"log"

	"github.com/google/uuid"
)

type Classification struct {
	TournamentName string
	Country        string
	Teams          []TeamClassification
}

type TeamClassification struct {
	TeamID         uuid.UUID
	TeamName       string
	Position       int
	Points         int
	GoalsFor       int
	GoalsAgainst   int
	GoalDifference int
}

func (a AppService) GetClassification(seasonID uuid.UUID) ([]Classification, error) {
	classification, err := a.classificationRepo.GetClassification(seasonID)
	if err != nil {
		return nil, err
	}

	tournament, err := a.tournamentRepo.GetTournamentBySeasonID(seasonID)
	if err != nil {
		log.Println("Error Get Tournament on GetClassification")
		return nil, err
	}

	teams := make([]TeamClassification, 0, len(classification))
	for _, c := range classification {
		teams = append(teams, TeamClassification{
			TeamID:         c.TeamID,
			TeamName:       c.TeamName,
			Position:       c.Position,
			Points:         c.Points,
			GoalsFor:       c.GoalsFor,
			GoalsAgainst:   c.GoalsAgainst,
			GoalDifference: c.GoalDifference,
		})
	}

	result := []Classification{
		{
			TournamentName: tournament.Name,
			Country:        tournament.CountryCode,
			Teams:          teams,
		},
	}

	return result, nil
}

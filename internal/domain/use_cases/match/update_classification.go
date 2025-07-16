package match

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

const (
	winPoints  = 3
	drawPoints = 1
	losePoints = 0
)

func (a AppService) UpdateClassification(homeTeamID, awayTeamID uuid.UUID, homeGoals, awayGoals int) error {
	var homeClassification, awayClassification domain.Classification
	var homePoints, awayPoints int

	switch {
	case homeGoals > awayGoals:
		homePoints = winPoints
		awayPoints = losePoints

	case homeGoals < awayGoals:
		homePoints = losePoints
		awayPoints = winPoints

	default:
		homePoints = drawPoints
		awayPoints = drawPoints
	}

	homeClassification.Points = homePoints
	awayClassification.Points = awayPoints

	homeClassification.GoalsFor = homeGoals
	awayClassification.GoalsFor = awayGoals

	homeClassification.GoalsAgainst = awayGoals
	awayClassification.GoalsAgainst = homeGoals

	homeClassification.TeamID = homeTeamID
	awayClassification.TeamID = awayTeamID

	err := a.classificationRepo.UpdateClassification(homeClassification)
	if err != nil {
		return fmt.Errorf("Home UpdateClassification failed: %w", err)
	}

	err = a.classificationRepo.UpdateClassification(awayClassification)
	if err != nil {
		return fmt.Errorf("Away UpdateClassification failed: %w", err)
	}
	return nil
}

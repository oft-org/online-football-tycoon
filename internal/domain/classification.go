package domain

import "github.com/google/uuid"

type Classification struct {
	TeamID         uuid.UUID
	TeamName       string
	Position       int
	Points         int
	GoalsFor       int
	GoalsAgainst   int
	GoalDifference int
}

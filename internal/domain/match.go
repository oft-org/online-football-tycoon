package domain

import (
	"time"

	"github.com/google/uuid"
)

type Match struct {
	HomeMatchStrategy Strategy
	AwayMatchStrategy Strategy
}

type SeasonMatch struct {
	ID         uuid.UUID
	SeasonID   uuid.UUID
	HomeTeamID uuid.UUID
	AwayTeamID uuid.UUID
	MatchDate  time.Time
	HomeResult *int
	AwayResult *int
}

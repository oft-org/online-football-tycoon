package domain

import (
	"time"

	"github.com/google/uuid"
)

type TournamentType string

const (
	TournamentLeague TournamentType = "League"
	TournamentCup    TournamentType = "Cup"
)

type Tournament struct {
	ID          uuid.UUID
	Name        string
	Type        TournamentType
	CountryCode string
	Division    int
	PromotionTo *uuid.UUID
	DescentTo   *uuid.UUID
}

type Season struct {
	ID           uuid.UUID
	TournamentID uuid.UUID
	FromDate     time.Time
	ToDate       time.Time
}

type SeasonTeam struct {
	SeasonID uuid.UUID
	TeamID   uuid.UUID
}

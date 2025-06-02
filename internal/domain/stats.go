package domain

import "github.com/google/uuid"

type Result struct {
	HomeStats TeamStats
	AwayStats TeamStats
}

type TeamStats struct {
	BallPossession int
	ScoringChances int
	Goals          int
}

type MatchEventStats struct {
	HomeEvents       []EventResult
	AwayEvents       []EventResult
	HomeScoreChances int
	AwayScoreChances int
	HomeGoals        int
	AwayGoals        int
}

type MatchEventInfo struct {
	MatchID     uuid.UUID
	TeamId      uuid.UUID
	EventType   string
	Minute      int
	Description string
}

type Event struct {
	Name    string
	Execute func() (string, int, int, int, int, error)
}

type EventResult struct {
	Event     string    `json:"event"`
	Minute    int       `json:"minute"`
	EventType string    `json:"eventtype"`
	TeamId    uuid.UUID `json:"teamid"`
	TeamName  string    `json:"team"`
}

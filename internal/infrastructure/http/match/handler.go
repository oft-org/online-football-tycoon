package match

import (
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type MatchApp interface {
	PlayMatch(seasonID, matchID uuid.UUID) (domain.Result, error)
	GetPendingMatches(timestamp time.Time) ([]domain.SeasonMatch, error)
	GetMatchDetailsByID(matchID uuid.UUID) (*MatchResponse, error)
	GetSeasonMatches(seasonID uuid.UUID) ([]domain.SeasonMatch, error)
}

type TeamApp interface {
	GenerateSeason(seasonID uuid.UUID, startDate time.Time) error
}

func NewHandler(matchApp MatchApp, teamApp TeamApp) Handler {
	return Handler{
		matchApp: matchApp,
		teamApp:  teamApp,
	}
}

type Handler struct {
	matchApp MatchApp
	teamApp  TeamApp
}

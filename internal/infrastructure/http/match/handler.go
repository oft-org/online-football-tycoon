package match

import (
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type App interface {
	PlayMatch(seasonID, matchID uuid.UUID) (domain.Result, error)
	GetPendingMatches(timestamp time.Time) ([]domain.SeasonMatch, error)
}

type TeamApp interface {
	GenerateRoundRobinSchedule(seasonID uuid.UUID) error
}

func NewHandler(app App, teamApp TeamApp) Handler {
	return Handler{
		app:     app,
		teamApp: teamApp,
	}
}

type Handler struct {
	app     App
	teamApp TeamApp
}

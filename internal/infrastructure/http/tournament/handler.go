package tournament

import (
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type App interface {
	GetTournamentsByCountry(country string) ([]domain.Tournament, error)
}

func NewHandler(app App) *Handler {
	return &Handler{
		app: app,
	}
}

type Handler struct {
	app App
}

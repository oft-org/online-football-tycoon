package player

import (
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type App interface {
	GeneratePlayer(country, position string) (domain.Player, error)
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}

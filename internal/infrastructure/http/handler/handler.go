package http

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/match"
)

type App interface {
	PlayMatch(matchID uuid.UUID) (match.Result, error)
}

func NewHandler(app match.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}

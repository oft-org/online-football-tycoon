package country

import (
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type App interface {
	GetCountries() ([]domain.Country, error)
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}

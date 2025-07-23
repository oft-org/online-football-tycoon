package classification

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain/use_cases/classification"
)

type App interface {
	GetClassification(seasonID uuid.UUID) ([]classification.ClassificationWithTournament, error)
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}

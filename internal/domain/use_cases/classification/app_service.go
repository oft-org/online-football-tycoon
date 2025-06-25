package classification

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type Repository interface {
	GetClassification(seasonID uuid.UUID) ([]domain.Classification, error)
}

func NewApp(repository Repository) AppService {
	return AppService{
		repo: repository,
	}
}

type AppService struct {
	repo Repository
}

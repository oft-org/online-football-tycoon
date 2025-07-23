package classification

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type ClassificationRepository interface {
	GetClassification(seasonID uuid.UUID) ([]domain.Classification, error)
}

type TournamentRepository interface {
	GetTournamentBySeasonID(seasonID uuid.UUID) (domain.Tournament, error)
}

func NewApp(classificationRepository ClassificationRepository, tournamentRepository TournamentRepository) AppService {
	return AppService{
		classificationRepo: classificationRepository,
		tournamentRepo:     tournamentRepository,
	}
}

type AppService struct {
	classificationRepo ClassificationRepository
	tournamentRepo     TournamentRepository
}

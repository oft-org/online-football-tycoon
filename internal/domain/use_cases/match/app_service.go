package match

import (
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type Repository interface {
	GetMatchStrategyById(matchId uuid.UUID) (*domain.Match, error)
	PostMatch(seasonId, homeTeamId, awayTeamId uuid.UUID, matchDate time.Time, homeGoals, awayGoals int) error
	PostMatchEvent(event domain.MatchEventInfo) error
	PostMatches(matches []domain.SeasonMatch) error
	GetPendingMatches(timestamp time.Time) ([]domain.SeasonMatch, error)
	UpdateMatch(seasonMatch domain.SeasonMatch) error
}

type ClassificationRepository interface {
	UpdateClassification(domain.Classification) error
}

func NewApp(repository Repository, classificationRepo ClassificationRepository) AppService {
	return AppService{
		repo:               repository,
		classificationRepo: classificationRepo,
		simulator:          NewSimulator(),
	}
}

type AppService struct {
	repo               Repository
	classificationRepo ClassificationRepository
	simulator          Simulator
}

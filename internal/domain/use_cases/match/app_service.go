package match

import (
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type Repository interface {
	GetMatchById(matchId uuid.UUID) (*domain.Match, error)
	PostMatch(seasonId, homeTeamId, awayTeamId uuid.UUID, matchDate time.Time, homeGoals, awayGoals int) error
	PostMatchEvent(event domain.MatchEventInfo) error
	PostMatches(matches []domain.SeasonMatch) error
	GetPendingMatches(timestamp time.Time) ([]domain.SeasonMatch, error)
}

func NewApp(repository Repository) AppService {
	return AppService{
		repo:      repository,
		simulator: NewSimulator(),
	}
}

type AppService struct {
	repo      Repository
	simulator Simulator
}

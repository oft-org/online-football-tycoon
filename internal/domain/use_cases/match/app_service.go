package match

import (
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type MatchRepository interface {
	GetMatchStrategyById(matchId uuid.UUID) (*domain.Match, error)
	PostMatch(seasonId, homeTeamId, awayTeamId uuid.UUID, matchDate time.Time, homeGoals, awayGoals int) error
	PostMatchEvent(event domain.MatchEventInfo) error
	PostMatches(matches []domain.SeasonMatch) error
	GetPendingMatches(timestamp time.Time) ([]domain.SeasonMatch, error)
	UpdateMatch(seasonMatch domain.SeasonMatch) error
	GetMatchByID(matchID uuid.UUID) (domain.SeasonMatch, error)
	GetMatchEvents(matchID uuid.UUID) ([]domain.MatchEventInfo, error)
	GetSeasonMatches(seasonID uuid.UUID) ([]domain.SeasonMatch, error)
}

type ClassificationRepository interface {
	UpdateClassification(domain.Classification) error
}

type TeamRepository interface {
	GetTeamByID(teamID uuid.UUID) (domain.Team, error)
}

func NewApp(matchRepo MatchRepository, classificationRepo ClassificationRepository, teamRepo TeamRepository) AppService {
	return AppService{
		matchRepo:          matchRepo,
		classificationRepo: classificationRepo,
		teamRepo:           teamRepo,
		simulator:          NewSimulator(),
	}
}

type AppService struct {
	matchRepo          MatchRepository
	classificationRepo ClassificationRepository
	teamRepo           TeamRepository
	simulator          Simulator
}

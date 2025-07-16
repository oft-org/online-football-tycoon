package match_test

import (
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockMatchRepository struct {
	mock.Mock
}

func (m *MockMatchRepository) GetMatchById(matchID uuid.UUID) (*domain.Match, error) {
	args := m.Called(matchID)

	match, ok := args.Get(0).(*domain.Match)
	if !ok && args.Get(0) != nil {
		panic("expected *Match from mock but got something else")
	}
	return match, args.Error(1)
}

func (m *MockMatchRepository) PostMatch(seasonId, homeTeamId, awayTeamId uuid.UUID, matchDate time.Time, homeGoals, awayGoals int) error {
	args := m.Called(seasonId, homeTeamId, awayTeamId, matchDate, homeGoals, awayGoals)
	return args.Error(0)
}

func (m *MockMatchRepository) PostMatchEvent(event domain.MatchEventInfo) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockMatchRepository) PostMatches(matches []domain.SeasonMatch) error {
	args := m.Called(matches)
	return args.Error(0)
}
func (m *MockMatchRepository) GetPendingMatches(timestamp time.Time) ([]domain.SeasonMatch, error) {
	args := m.Called(timestamp)
	return args.Get(0).([]domain.SeasonMatch), args.Error(1)
}

type MockClassificationRepository struct {
	mock.Mock
}

func (m *MockClassificationRepository) UpdateClassification(classification domain.Classification) error {
	args := m.Called(classification)
	return args.Error(0)
}

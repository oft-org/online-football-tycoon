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

func (m *MockMatchRepository) GetMatchStrategyById(matchID uuid.UUID) (*domain.Match, error) {
	args := m.Called(matchID)
	match, _ := args.Get(0).(*domain.Match)
	return match, args.Error(1)
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

func (m *MockMatchRepository) UpdateMatch(seasonMatch domain.SeasonMatch) error {
	args := m.Called(seasonMatch.ID, seasonMatch.SeasonID, seasonMatch.HomeTeamID, seasonMatch.AwayTeamID, seasonMatch.MatchDate, seasonMatch.HomeResult, seasonMatch.AwayResult)
	return args.Error(0)
}

type MockTeamRepository struct {
	mock.Mock
}

func (m *MockTeamRepository) GetTeamByID(teamID uuid.UUID) (domain.Team, error) {
	args := m.Called(teamID)
	return args.Get(0).(domain.Team), args.Error(1)
}

type MockClassificationRepository struct {
	mock.Mock
}

func (m *MockClassificationRepository) UpdateClassification(classification domain.Classification) error {
	args := m.Called(classification)
	return args.Error(0)
}

func (m *MockMatchRepository) GetMatchByID(matchID uuid.UUID) (domain.SeasonMatch, error) {
	args := m.Called(matchID)
	return args.Get(0).(domain.SeasonMatch), args.Error(1)
}

func (m *MockMatchRepository) GetMatchEvents(matchID uuid.UUID) ([]domain.MatchEventInfo, error) {
	args := m.Called(matchID)
	return args.Get(0).([]domain.MatchEventInfo), args.Error(1)
}

func (m *MockMatchRepository) GetSeasonMatches(seasonID uuid.UUID) ([]domain.SeasonMatch, error) {
	args := m.Called(seasonID)
	return args.Get(0).([]domain.SeasonMatch), args.Error(1)
}

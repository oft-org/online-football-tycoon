package match

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (a AppService) PlayMatch(matchID uuid.UUID) (domain.Result, error) {
	m, err := a.repo.GetMatchById(matchID)
	if err != nil {
		return domain.Result{}, fmt.Errorf("error retrieving match: %w", err)
	}
	if m == nil {
		log.Printf("repo.GetMatchById returned nil for matchID: %s", matchID)
		return domain.Result{}, fmt.Errorf("no match found with ID: %s", matchID)
	}
	result, allEvents, err := a.simulator.Play(m)
	if err != nil {
		return domain.Result{}, fmt.Errorf("error playing match: %w", err)
	}

	matchDate := time.Now()
	homeTeamId := m.HomeMatchStrategy.StrategyTeam.Id
	awayTeamId := m.AwayMatchStrategy.StrategyTeam.Id

	log.Printf("Match played: homeTeamID=%s, awayTeamID=%s", homeTeamId, awayTeamId)

	err = a.repo.PostMatch(homeTeamId, awayTeamId, matchDate, result.HomeStats.Goals, result.AwayStats.Goals)
	if err != nil {
		return domain.Result{}, fmt.Errorf("error PostMatch: %w", err)
	}

	for _, event := range allEvents {

		matchEventInfo := domain.MatchEventInfo{
			MatchID:     matchID,
			TeamId:      event.TeamId,
			EventType:   event.EventType,
			Minute:      event.Minute,
			Description: event.Event,
		}

		err = a.repo.PostMatchEvent(matchEventInfo)
		if err != nil {
			log.Printf("error posting event to repo: %v", err)
			return domain.Result{}, fmt.Errorf("PostMatchEvent failed: %w", err)
		}
	}
	return result, nil
}

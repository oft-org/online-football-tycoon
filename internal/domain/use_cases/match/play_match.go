package match

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (a AppService) PlayMatch(seasonID, matchID uuid.UUID) (domain.Result, error) {
	m, err := a.repo.GetMatchStrategyById(matchID)
	if err != nil {
		return domain.Result{}, fmt.Errorf("error retrieving match: %w", err)
	}
	if m == nil {
		log.Printf("repo.GetMatchStrategyById returned nil for matchID: %s", matchID)
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

	var seasonMatch domain.SeasonMatch
	seasonMatch.ID = matchID
	seasonMatch.SeasonID = seasonID
	seasonMatch.HomeTeamID = homeTeamId
	seasonMatch.AwayTeamID = awayTeamId
	seasonMatch.MatchDate = matchDate
	seasonMatch.HomeResult = &result.HomeStats.Goals
	seasonMatch.AwayResult = &result.AwayStats.Goals

	log.Printf("Calling UpdateMatch with HomeResult=%v, AwayResult=%v", seasonMatch.HomeResult, seasonMatch.AwayResult)
	err = a.repo.UpdateMatch(seasonMatch)
	if err != nil {
		return domain.Result{}, fmt.Errorf("error UpdateMatch: %w", err)
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

	err = a.UpdateClassification(homeTeamId, awayTeamId, result.HomeStats.Goals, result.AwayStats.Goals)
	if err != nil {
		log.Printf("error posting event to repo: %v", err)
		return domain.Result{}, fmt.Errorf("UpdateClassification failed: %w", err)
	}

	return result, nil
}

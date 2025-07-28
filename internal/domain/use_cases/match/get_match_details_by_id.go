package match

import (
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
	httpMatch "github.com/robertobouses/online-football-tycoon/internal/infrastructure/http/match"
)

func (a *AppService) GetMatchDetailsByID(matchID uuid.UUID) (*httpMatch.MatchResponse, error) {
	match, err := a.matchRepo.GetMatchByID(matchID)
	if err != nil {
		log.Printf("Error getting match by ID: %v", err)
		return nil, err
	}
	log.Printf("GetMatchByID returned match: ID=%s, HomeResult=%v, AwayResult=%v", match.ID, match.HomeResult, match.AwayResult)

	homeTeam, err := a.teamRepo.GetTeamByID(match.HomeTeamID)
	if err != nil {
		log.Printf("Error getting home team: %v", err)
		return nil, err
	}

	awayTeam, err := a.teamRepo.GetTeamByID(match.AwayTeamID)
	if err != nil {
		log.Printf("Error getting away team: %v", err)
		return nil, err
	}

	var events []domain.MatchEventInfo
	if match.HomeResult != nil && match.AwayResult != nil {
		events, err = a.matchRepo.GetMatchEvents(matchID)
		if err != nil {
			log.Printf("Error getting match events: %v", err)
			return nil, err
		}
		log.Printf("Number of events: %d", len(events))
	}

	httpEvents := make([]httpMatch.MatchEvent, len(events))
	for i, e := range events {
		httpEvents[i] = httpMatch.MatchEvent{
			ID:        e.ID,
			EventType: e.EventType,
			Minute:    e.Minute,
		}
	}

	return &httpMatch.MatchResponse{
		MatchID:   match.ID,
		MatchDate: match.MatchDate,
		HomeTeam: httpMatch.TeamInfo{
			ID:   match.HomeTeamID,
			Name: homeTeam.Name,
		},
		AwayTeam: httpMatch.TeamInfo{
			ID:   match.AwayTeamID,
			Name: awayTeam.Name,
		},
		HomeResult: match.HomeResult,
		AwayResult: match.AwayResult,
		Events:     httpEvents,
	}, nil
}

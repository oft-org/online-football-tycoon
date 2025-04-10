package repository

import (
	"encoding/json"
	"log"

	"github.com/robertobouses/online-football-tycoon/match"
)

func (r *repository) PostMatchEvent(matchEvent match.MatchEvent) error {
	homeEventsJSON, err := json.Marshal(matchEvent.MatchInfo.HomeEvents)
	if err != nil {
		log.Println("Error marshalling home events:", err)
		return err
	}

	awayEventsJSON, err := json.Marshal(matchEvent.MatchInfo.AwayEvents)
	if err != nil {
		log.Println("Error marshalling away events:", err)
		return err
	}

	_, err = r.postMatchEvents.Exec(
		matchEvent.MatchID,
		string(homeEventsJSON),
		string(awayEventsJSON),
		matchEvent.MatchInfo.HomeScoreChances,
		matchEvent.MatchInfo.AwayScoreChances,
		matchEvent.MatchInfo.HomeGoals,
		matchEvent.MatchInfo.AwayGoals,
	)

	if err != nil {
		log.Print("Error executing PostMatch statement:", err)
		return err
	}

	return nil
}

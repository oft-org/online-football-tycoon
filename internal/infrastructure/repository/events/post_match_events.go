package repository

import (
	"log"

	"github.com/robertobouses/online-football-tycoon/match"
)

func (r *repository) PostMatchEvent(matchEventInfo match.MatchEventInfo) error {

	_, err := r.postMatchEvents.Exec(
		matchEventInfo.MatchID,
		matchEventInfo.TeamId,
		matchEventInfo.EventType,
		matchEventInfo.Minute,
		matchEventInfo.Description,
	)

	if err != nil {
		log.Print("Error executing PostMatchEvent statement:", err)
		return err
	}

	return nil
}

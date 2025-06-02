package match

import (
	"log"

	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) PostMatchEvent(matchEventInfo domain.MatchEventInfo) error {

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

package match

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) GetMatchEvents(matchID uuid.UUID) ([]domain.MatchEventInfo, error) {
	rows, err := r.getMatchEvents.Query(matchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []domain.MatchEventInfo

	for rows.Next() {
		var matchEventInfo domain.MatchEventInfo

		err := rows.Scan(
			&matchEventInfo.ID,
			&matchEventInfo.MatchID,
			&matchEventInfo.TeamId,
			&matchEventInfo.EventType,
			&matchEventInfo.Minute,
			&matchEventInfo.Description,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, matchEventInfo)
	}
	return events, nil
}

package match

import (
	"time"

	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (a AppService) GetPendingMatches(timestamp time.Time) ([]domain.SeasonMatch, error) {
	matches, err := a.repo.GetPendingMatches(timestamp)
	if err != nil {
		return []domain.SeasonMatch{}, err
	}
	return matches, nil

}

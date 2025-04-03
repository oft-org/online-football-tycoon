package match

import (
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	GetMatchById(matchId uuid.UUID) (*Match, error)
	PostMatch(homeTeamId, awayTeamId uuid.UUID, matchDate time.Time, homeGoals, awayGoals int) error
}

func NewApp(repository Repository) AppService {
	return AppService{repo: repository}
}

type AppService struct {
	repo Repository
}

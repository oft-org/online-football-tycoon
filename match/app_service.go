package match

import "github.com/google/uuid"

type Repository interface {
	GetMatchById(matchId uuid.UUID) (*Match, error)
}

func NewApp(repository Repository) AppService {
	return AppService{repo: repository}
}

type AppService struct {
	repo Repository
}

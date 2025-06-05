package player

import "github.com/robertobouses/online-football-tycoon/internal/domain"

type Repository interface {
	PostPlayer(player domain.Player) error
}

func NewApp(repository Repository) AppService {
	return AppService{
		repo: repository,
	}
}

type AppService struct {
	repo Repository
}

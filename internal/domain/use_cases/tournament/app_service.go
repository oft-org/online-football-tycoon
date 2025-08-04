package tournament

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type TournamentRepository interface {
	GetTournamentsByCountry(country string) ([]domain.Tournament, error)
	GetTournamentBySeasonID(seasonId uuid.UUID) (domain.Tournament, error)
}

type TeamRepository interface {
	GetSeasonTeam(seasonID uuid.UUID) ([]uuid.UUID, error)
}

type MatchRepository interface {
	PostMatches(matches []domain.SeasonMatch) error
}

func NewApp(tournamentRepository TournamentRepository, teamRepository TeamRepository, matchRepository MatchRepository) AppService {
	return AppService{
		tournamentRepo: tournamentRepository,
		teamRepo:       teamRepository,
		matchRepo:      matchRepository,
	}
}

type AppService struct {
	tournamentRepo TournamentRepository
	teamRepo       TeamRepository
	matchRepo      MatchRepository
}

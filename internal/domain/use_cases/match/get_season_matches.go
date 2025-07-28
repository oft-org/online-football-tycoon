package match

import (
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (a AppService) GetSeasonMatches(seasonID uuid.UUID) ([]domain.SeasonMatch, error) {
	seasonMatches, err := a.matchRepo.GetSeasonMatches(seasonID)
	if err != nil {
		log.Println("Error GetSesaonMatches from repo")
		return []domain.SeasonMatch{}, err
	}
	return seasonMatches, nil

}

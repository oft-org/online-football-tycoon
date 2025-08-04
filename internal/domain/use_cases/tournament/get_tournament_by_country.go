package tournament

import (
	"log"

	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (a AppService) GetTournamentsByCountry(country string) ([]domain.Tournament, error) {

	tournaments, err := a.tournamentRepo.GetTournamentsByCountry(country)
	if err != nil {
		log.Println("Error Get Tournament on GetTournamentsByCountry")
		return nil, err
	}

	return tournaments, nil
}

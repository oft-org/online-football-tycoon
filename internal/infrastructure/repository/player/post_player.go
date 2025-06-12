package match

import (
	"log"

	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) PostPlayer(player domain.Player) error {
	_, err := r.postPlayer.Exec(
		player.FirstName,
		player.LastName,
		player.Nationality,
		player.Position,
		player.Age,
		player.Fee,
		player.Salary,
		player.Technique,
		player.Mental,
		player.Physique,
		player.InjuryDays,
		player.Lined,
		player.Familiarity,
		player.Fitness,
		player.Happiness,
	)

	if err != nil {
		log.Print("Error executing PostPlayer statement:", err)
		return err
	}

	return nil
}

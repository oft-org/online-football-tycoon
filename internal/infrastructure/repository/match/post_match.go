package match

import (
	"log"
	"time"

	"github.com/google/uuid"
)

// TODO: Esto no te va a funcionar (creo) estás definiendo un método para un struct que está en otro package.
// Aunque funcionase, que lo dudo, no lo hagas.

//DONE: Está en verde... No comprendo la pregunta, si te refieres a que el struct es Repository esá en este mismo paquete...

func (r *Repository) PostMatch(homeTeamId, awayTeamId uuid.UUID, matchDate time.Time, homeGoals, awayGoals int) error {
	_, err := r.postMatch.Exec(
		homeTeamId,
		awayTeamId,
		matchDate,
		homeGoals,
		awayGoals,
	)

	if err != nil {
		log.Print("Error executing PostMatch statement:", err)
		return err
	}

	return nil
}

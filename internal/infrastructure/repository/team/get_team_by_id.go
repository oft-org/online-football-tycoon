package team

import (
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) GetTeamByID(teamID uuid.UUID) (domain.Team, error) {
	row := r.getTeamByID.QueryRow(teamID)

	var team domain.Team
	err := row.Scan(
		&team.Id,
		&team.Name,
		&team.Country,
	)
	if err != nil {
		return domain.Team{}, err
	}

	log.Printf("GetTeamByID returned team: ID=%v, Name=%v, Country=%v",
		team.Id, team.Name, team.Country)

	return team, nil
}

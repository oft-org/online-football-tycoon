package tournament

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) GetTournamentBySeasonID(seasonId uuid.UUID) (*domain.Tournament, error) {
	row := r.getTournamentBySeasonID.QueryRow(seasonId)
	var tournament domain.Tournament
	if err := row.Scan(
		&tournament.ID,
		&tournament.Name,
		&tournament.Type,
		&tournament.CountryCode,
		&tournament.Division,
		&tournament.PromotionTo,
		&tournament.DescentTo,
	); err != nil {
		return &domain.Tournament{}, err
	}
	return &tournament, nil
}

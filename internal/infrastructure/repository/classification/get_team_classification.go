package classification

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) GetTeamClassification(teamID uuid.UUID) (*domain.Classification, error) {
	row := r.getTeamClassification.QueryRow(teamID)
	var classification domain.Classification
	if err := row.Scan(
		&classification.Points,
		&classification.GoalsFor,
		&classification.GoalsAgainst,
	); err != nil {
		return nil, err
	}
	return &classification, nil
}

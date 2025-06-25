package classification

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) GetClassification(seasonID uuid.UUID) ([]domain.Classification, error) {
	rows, err := r.getClassification.Query(seasonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clasification []domain.Classification
	for rows.Next() {
		var classified domain.Classification

		err := rows.Scan(
			&classified.TeamID,
			&classified.TeamName,
			&classified.Position,
			&classified.Points,
			&classified.GoalsFor,
			&classified.GoalsAgainst,
			&classified.GoalDifference,
		)
		if err != nil {
			return nil, err
		}
		clasification = append(clasification, classified)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clasification, nil
}

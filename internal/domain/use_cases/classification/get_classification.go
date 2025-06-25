package classification

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (a AppService) GetClassification(seasonID uuid.UUID) ([]domain.Classification, error) {
	classification, err := a.repo.GetClassification(seasonID)
	if err != nil {
		return []domain.Classification{}, err
	}
	return classification, nil

}

package classification

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ClassificationInfo struct {
	TeamID         uuid.UUID `json:"team_id"`
	TeamName       string    `json:"team_name"`
	Position       int       `json:"position"`
	Points         int       `json:"points"`
	GoalsFor       int       `json:"goals_for"`
	GoalsAgainst   int       `json:"goals_against"`
	GoalDifference int       `json:"goal_difference"`
	TournamentName string    `json:"tournament_name"`
	Country        string    `json:"country"`
}

func (h *Handler) GetClassification(c *gin.Context) {
	seasonIDParam := c.Param("season_id")
	seasonID, err := uuid.Parse(seasonIDParam)
	if err != nil {
		log.Printf("Invalid season_id: %s | Error: %v", seasonIDParam, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid season_id"})
		return
	}

	classificationData, err := h.app.GetClassification(seasonID)
	if err != nil {
		log.Printf("Failed to get classification for season_id %s | Error: %v", seasonID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get classification"})
		return
	}

	var response []ClassificationInfo
	for _, c := range classificationData {
		response = append(response, ClassificationInfo{
			TeamID:         c.TeamID,
			TeamName:       c.TeamName,
			Position:       c.Position,
			Points:         c.Points,
			GoalsFor:       c.GoalsFor,
			GoalsAgainst:   c.GoalsAgainst,
			GoalDifference: c.GoalDifference,
			TournamentName: c.TournamentName,
			Country:        c.Country,
		})
	}

	c.JSON(http.StatusOK, response)
}

package classification

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

	c.JSON(http.StatusOK, classificationData)
}

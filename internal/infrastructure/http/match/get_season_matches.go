package match

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetSeasonMatches(c *gin.Context) {
	seasonIDString := c.Query("season_id")
	if seasonIDString == "" {
		log.Println("Missing 'seasonID' query parameter")
		c.JSON(http.StatusBadRequest, gin.H{"error": "seasonID query param is required"})
		return
	}

	seasonID, err := uuid.Parse(seasonIDString)
	if err != nil {
		log.Println("Error parse seasonID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Parse seasonID"})
	}

	seasonMatches, err := h.matchApp.GetSeasonMatches(seasonID)
	if err != nil {
		log.Printf("Failed to get season matches with season_id: %s | Error: %v", seasonID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get season matches"})
		return
	}

	c.JSON(http.StatusOK, seasonMatches)
}

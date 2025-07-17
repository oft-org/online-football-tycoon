package match

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPendingMatches(c *gin.Context) {
	timestampStr := c.Query("timestamp")
	if timestampStr == "" {
		log.Println("Missing 'timestamp' query parameter")
		c.JSON(http.StatusBadRequest, gin.H{"error": "timestamp query param is required"})
		return
	}

	timestamp, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		log.Printf("Invalid timestamp format: %s | Error: %v", timestampStr, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid timestamp format, must be RFC3339"})
		return
	}

	pendingMatches, err := h.app.GetPendingMatches(timestamp)
	if err != nil {
		log.Printf("Failed to get pending matches after %s | Error: %v", timestamp, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get pending matches"})
		return
	}

	c.JSON(http.StatusOK, pendingMatches)
}

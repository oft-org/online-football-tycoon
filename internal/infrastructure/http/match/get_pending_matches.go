package match

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPendingMatches(c *gin.Context) {
	timestampStr := c.Query("timestamp")
	if timestampStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "timestamp query param is required"})
		return
	}

	timestamp, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid timestamp format, must be RFC3339"})
		return
	}

	pendingMatches, err := h.app.GetPendingMatches(timestamp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get pending matches"})
		return
	}

	c.JSON(http.StatusOK, pendingMatches)
}

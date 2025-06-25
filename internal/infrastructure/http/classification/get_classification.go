package classification

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetClassification(c *gin.Context) {
	seasonIDParam := c.Param("season_id")
	seasonID, err := uuid.Parse(seasonIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid season_id"})
		return
	}

	classificationData, err := h.app.GetClassification(seasonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get classification"})
		return
	}

	c.JSON(http.StatusOK, classificationData)
}

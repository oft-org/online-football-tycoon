package match

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoundRobinScheduleRequest struct {
	SeasonID uuid.UUID `json:"season_id"`
}

func (h Handler) PostSeasonMatches(c *gin.Context) {
	var req RoundRobinScheduleRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostSeasonMatches] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.teamApp.GenerateRoundRobinSchedule(req.SeasonID)
	if err != nil {
		log.Printf("[PostSeasonMatches] error generating schedule for season %s: %v", req.SeasonID, err)
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{
		"mensaje": "Season created",
	})
}

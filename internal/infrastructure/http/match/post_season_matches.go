package match

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoundRobinScheduleRequest struct {
	SeasonID  uuid.UUID `json:"season_id"`
	StartDate string    `json:"start_date"`
}

func (h Handler) PostSeasonMatches(c *gin.Context) {
	var req RoundRobinScheduleRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostSeasonMatches] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		log.Printf("[PostSeasonMatches] invalid date format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format. Use YYYY-MM-DD"})
		return
	}

	err = h.tournamentApp.GenerateSeason(req.SeasonID, startDate)
	if err != nil {
		log.Printf("[GenerateSeason] error generating schedule for season %s: %v", req.SeasonID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Season matches generated successfully",
	})
}

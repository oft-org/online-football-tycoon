package match

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TeamInfo struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type MatchEvent struct {
	ID        uuid.UUID `json:"id"`
	EventType string    `json:"event_type"`
	Minute    int       `json:"minute"`
}

type MatchResponse struct {
	MatchID    uuid.UUID    `json:"match_id"`
	MatchDate  time.Time    `json:"match_date"`
	HomeTeam   TeamInfo     `json:"home_team"`
	AwayTeam   TeamInfo     `json:"away_team"`
	HomeResult *int         `json:"home_result,omitempty"`
	AwayResult *int         `json:"away_result,omitempty"`
	Events     []MatchEvent `json:"events,omitempty"`
}

func (h *Handler) GetMatchByID(c *gin.Context) {
	matchIDString := c.Param("match_id")
	matchID, err := uuid.Parse(matchIDString)
	if err != nil {
		log.Printf("Invalid match_id: %s | Error: %v", matchIDString, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid match_id"})
		return
	}

	if matchID == uuid.Nil {
		log.Println("Missing 'match_id' query parameter")
		c.JSON(http.StatusBadRequest, gin.H{"error": "match_id query param is required"})
		return
	}

	resp, err := h.matchApp.GetMatchDetailsByID(matchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get match details"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

package match

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MatchRequest struct {
	SeasonId uuid.UUID `json:"season_id"`
	MatchId  uuid.UUID `json:"match_id"`
}

func (h Handler) PostPlayMatchbyId(c *gin.Context) {
	var req MatchRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostPlayMatchbyId] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.Printf("match id: %s", req.MatchId)

	result, err := h.matchApp.PlayMatch(req.SeasonId, req.MatchId)
	if err != nil {
		log.Printf("[PostPlayMatchbyId] error playing match %s: %v", req.MatchId, err)
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{
		"mensaje":   "Encuentro simulado",
		"resultado": result,
	})
}

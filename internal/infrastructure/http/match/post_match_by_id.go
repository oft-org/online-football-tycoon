package match

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MatchId struct {
	MatchId uuid.UUID `json:"match_id"`
}

func (h Handler) PostMatchbyId(c *gin.Context) {
	var req MatchId
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostMatchbyId] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.Printf("match id: %s", req.MatchId)

	result, err := h.app.PlayMatch(req.MatchId)
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{
		"mensaje":   "Encuentro simulado",
		"resultado": result,
	})
}

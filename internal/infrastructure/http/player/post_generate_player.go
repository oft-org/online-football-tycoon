package player

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

type PostGeneratePlayerRequest struct {
	Country  string `json:"country"`
	Position string `json:"position"`
}

func (h Handler) PostGeneratePlayer(c *gin.Context) {
	var req PostGeneratePlayerRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostGeneratePlayerRequest] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	player, err := h.app.GeneratePlayer(req.Country, req.Position)
	if err != nil {
		log.Printf("[PostGeneratePlayer] error generating player (country=%s, position=%s): %v", req.Country, req.Position, err)
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{
		"mensaje":            "Jugadores generados correctamente",
		"jugadores_generado": player,
	})
}

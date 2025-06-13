package match

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) PostSeasonMatches(c *gin.Context) {

	err := h.teamApp.GenerateRoundRobinSchedule()
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{
		"mensaje": "Season created",
	})
}

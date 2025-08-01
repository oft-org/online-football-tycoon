package tournament

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTournamentsByCountry(c *gin.Context) {
	country := c.Param("country")

	if country == "" {
		log.Println("[WARN] Missing country parameter")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing country parameter"})
		return
	}

	log.Printf("[INFO] Getting tournaments for country: %s", country)

	tournaments, err := h.app.GetTournamentsByCountry(country)
	if err != nil {
		log.Printf("[ERROR] Failed to get tournaments for country %s: %v", country, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tournaments"})
		return
	}
	if len(tournaments) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("no tournaments found for country '%s'", country),
			"data":    []any{},
		})
		return
	}
	log.Printf("[INFO] Found %d tournaments for country %s", len(tournaments), country)
	c.JSON(http.StatusOK, tournaments)
}

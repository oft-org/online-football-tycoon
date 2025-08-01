package country

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCountries(c *gin.Context) {

	countries, err := h.app.GetCountries()
	if err != nil {
		log.Printf("Failed to get countries | Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get countries"})
		return
	}

	c.JSON(http.StatusOK, countries)
}

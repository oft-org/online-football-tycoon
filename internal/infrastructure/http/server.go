package http

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/online-football-tycoon/internal/infrastructure/http/classification"
	"github.com/robertobouses/online-football-tycoon/internal/infrastructure/http/country"
	"github.com/robertobouses/online-football-tycoon/internal/infrastructure/http/match"
	"github.com/robertobouses/online-football-tycoon/internal/infrastructure/http/player"
	"github.com/robertobouses/online-football-tycoon/internal/infrastructure/http/tournament"
)

type Server struct {
	match          match.Handler
	player         player.Handler
	classification classification.Handler
	country        country.Handler
	tournament     tournament.Handler
	engine         *gin.Engine
}

func NewServer(
	match match.Handler,
	player player.Handler,
	classification classification.Handler,
	country country.Handler,
	tournament tournament.Handler,

) Server {

	return Server{
		match:          match,
		player:         player,
		classification: classification,
		country:        country,
		tournament:     tournament,
		engine:         gin.Default(),
	}
}

func (s *Server) Run(port string) error {
	s.engine.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET, PUT, POST, DELETE, PATCH, OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "X-Accept-Language"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))

	match := s.engine.Group("/match")
	match.POST("/play", s.match.PostPlayMatchbyId)
	match.POST("/season", s.match.PostSeasonMatches)
	match.GET("/pending", s.match.GetPendingMatches)
	match.GET("/:match_id", s.match.GetMatchByID)
	match.GET("/season", s.match.GetSeasonMatches)

	player := s.engine.Group("/player")
	player.POST("/generate", s.player.PostGeneratePlayer)

	classification := s.engine.Group("/season")
	classification.GET("/:season_id/classification", s.classification.GetClassification)

	country := s.engine.Group("/country")
	country.GET("/", s.country.GetCountries)

	tournament := s.engine.Group("/tournament")
	tournament.GET("/:country", s.tournament.GetTournamentsByCountry)

	log.Printf("running api at %s port\n", port)
	return s.engine.Run(fmt.Sprintf(":%s", port))
}

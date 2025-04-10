package match

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/team"
)

func CalculateNumberOfMatchEvents(homeGameTempo, awayGameTempo string) (int, error) {

	var tempoMap = map[string]int{
		"slow_tempo":     1,
		"balanced_tempo": 2,
		"fast_tempo":     3,
	}

	homeTempo := tempoMap[homeGameTempo]
	awayTempo := tempoMap[awayGameTempo]
	matchTempo := homeTempo + awayTempo
	var numberOfMatchEvents int

	switch {
	case matchTempo <= 2:
		numberOfMatchEvents = rand.Intn(6) + 3
	case matchTempo > 2 && matchTempo <= 3:
		numberOfMatchEvents = rand.Intn(8) + 4
	case matchTempo > 3 && matchTempo <= 4:
		numberOfMatchEvents = rand.Intn(9) + 6
	case matchTempo > 4 && matchTempo <= 5:
		numberOfMatchEvents = rand.Intn(9) + 9
	case matchTempo > 5 && matchTempo <= 6:
		numberOfMatchEvents = rand.Intn(11) + 12
	}

	log.Println("numberOfMatchEvents", numberOfMatchEvents)
	return numberOfMatchEvents, nil
}

func DistributeMatchEvents(home, awayHome team.Team, numberOfMatchEvents int) (int, int, error) {
	log.Println("home en DistributeMatchEvents", home)
	log.Println("away en DistributeMatchEvents", awayHome)
	homeTotalQuality, err := CalculateQuality(home)
	if err != nil {
		return 0, 0, err
	}
	log.Println("total home Quality", homeTotalQuality)
	awayTotalQuality, err := CalculateQuality(awayHome)
	if err != nil {
		return 0, 0, err
	}
	log.Println("total away Quality", awayTotalQuality)
	allQuality := homeTotalQuality + awayTotalQuality
	var homeEvents int
	homeProportion := float64(homeTotalQuality) / float64(allQuality)

	homeEvents = int(homeProportion*float64(numberOfMatchEvents)) + rand.Intn(4) + 2

	log.Printf("number of home events %v ANTES DE RANDOMFACTOR", homeEvents)

	randomFactor := rand.Intn(11) - 5

	homeEvents += randomFactor

	awayEvents := numberOfMatchEvents - homeEvents
	log.Printf("number of home events %v, away events %v Despues DE RANDOMFACTOR", homeEvents, awayEvents)
	if homeEvents <= 0 {
		homeEvents = 0
	}
	if awayEvents < 0 {
		awayEvents = 0
	}
	log.Printf("number of home events %v, away events %v", homeEvents, awayEvents)
	return homeEvents, awayEvents, nil
}

func CalculateQuality(home team.Team) (int, error) {
	var totalTechnique, totalMental, totalPhysique int
	for _, player := range home.Players {
		totalTechnique += player.Technique
		totalMental += player.Mental
		totalPhysique += player.Physique
	}

	return 2*totalTechnique + 3*totalMental + 2*totalPhysique, nil
}

func clamp(value int, min int, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

func GetRandomDefender(home []team.Player) *team.Player {
	var defenders []team.Player
	for _, player := range home {
		if player.Position == "defender" {
			defenders = append(defenders, player)
		}

	}
	return GetRandomPlayer(defenders)
}

func GetRandomMidfielder(home []team.Player) *team.Player {
	var midfielders []team.Player
	for _, player := range home {
		if player.Position == "midfielder" {
			midfielders = append(midfielders, player)
		}

	}
	return GetRandomPlayer(midfielders)
}

func GetRandomForward(home []team.Player) *team.Player {
	var forwards []team.Player
	for _, player := range home {
		if player.Position == "forward" {
			forwards = append(forwards, player)
		}

	}
	return GetRandomPlayer(forwards)
}

func GetGoalkeeper(home []team.Player) *team.Player {
	var goalkeepers []team.Player
	for _, player := range home {
		if player.Position == "goalkeeper" {
			goalkeepers = append(goalkeepers, player)
		}

	}
	return GetRandomPlayer(goalkeepers)
}

func GetRandomPlayerExcludingGoalkeeper(home []team.Player) *team.Player {
	var playersExcludingGoalkeepers []team.Player
	for _, player := range home {
		if player.Position != "goalkeeper" {
			playersExcludingGoalkeepers = append(playersExcludingGoalkeepers, player)
		}

	}
	return GetRandomPlayer(playersExcludingGoalkeepers)
}

func GetRandomPlayer(filteredPlayers []team.Player) *team.Player {
	if len(filteredPlayers) == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	randomPlayer := filteredPlayers[rand.Intn(len(filteredPlayers))]
	return &randomPlayer
}

type Event struct {
	Name    string
	Execute func() (string, int, int, int, int, error)
}

type EventResult struct {
	Event     string    `json:"event"`
	Minute    int       `json:"minute"`
	EventType string    `json:"eventtype"`
	TeamId    uuid.UUID `json:"teamid"`
	TeamName  string    `json:"team"`
}

func GenerateEvents(home, awayHome team.Team, numberOfHomeEvents, numberOfAwayEvents int) MatchEventStats {

	homeEvents := []Event{
		{
			"Pase clave",
			func() (string, int, int, int, int, error) {
				return KeyPass(home, awayHome)
			},
		},
		{
			"Remate a puerta",
			func() (string, int, int, int, int, error) {
				return Shot(home, awayHome, GetRandomForward(home.Players))
			},
		},
		{
			"Penalty",
			func() (string, int, int, int, int, error) {
				return PenaltyKick(home, awayHome)
			},
		},
		{
			"Tiro lejano",
			func() (string, int, int, int, int, error) {
				return LongShot(home, awayHome)
			},
		},
		{
			" Lanzamiento de Falta Indirecta",
			func() (string, int, int, int, int, error) {
				return IndirectFreeKick(home, awayHome)
			},
		},
		{
			"Regate",
			func() (string, int, int, int, int, error) {
				return Dribble(home, awayHome)
			},
		},
		{
			"Falta",
			func() (string, int, int, int, int, error) {
				return Foul(home, awayHome, nil)
			},
		},

		{
			"Gran Ocasión",
			func() (string, int, int, int, int, error) {
				return GreatScoringChance(home)
			},
		},
		{
			"Córner",
			func() (string, int, int, int, int, error) {
				return CornerKick(home, awayHome)
			},
		},
		{
			"Fuera de Juego",
			func() (string, int, int, int, int, error) {
				return Offside(home, awayHome)
			},
		},
		{
			"Cabezazo",
			func() (string, int, int, int, int, error) {
				return Headed(home, awayHome)
			},
		}, {
			"Contragolpe",
			func() (string, int, int, int, int, error) {
				return CounterAttack(home, awayHome)
			},
		},
	}

	awayEvents := []Event{
		{
			"Pase clave",
			func() (string, int, int, int, int, error) {
				return KeyPass(awayHome, home)
			},
		},
		{
			"Remate a puerta",
			func() (string, int, int, int, int, error) {
				return Shot(awayHome, home, GetRandomForward(awayHome.Players))
			},
		},
		{
			"Penalty",
			func() (string, int, int, int, int, error) {
				return PenaltyKick(awayHome, home)
			},
		},
		{
			"Tiro lejano",
			func() (string, int, int, int, int, error) {
				return LongShot(awayHome, home)
			},
		},
		{
			" Lanzamiento de Falta Indirecta",
			func() (string, int, int, int, int, error) {
				return IndirectFreeKick(awayHome, home)
			},
		},
		{
			"Regate",
			func() (string, int, int, int, int, error) {
				return Dribble(awayHome, home)
			},
		},
		{
			"Falta",
			func() (string, int, int, int, int, error) {
				return Foul(awayHome, home, nil)
			},
		},
		{
			"Gran Ocasión",
			func() (string, int, int, int, int, error) {
				return GreatScoringChance(awayHome)
			},
		},
		{
			"Córner",
			func() (string, int, int, int, int, error) {
				return CornerKick(awayHome, home)
			},
		},
		{
			"Fuera de Juego",
			func() (string, int, int, int, int, error) {
				return Offside(awayHome, home)
			},
		},
		{
			"Cabezazo",
			func() (string, int, int, int, int, error) {
				return Headed(awayHome, home)
			},
		}, {
			"Contragolpe",
			func() (string, int, int, int, int, error) {
				return CounterAttack(awayHome, home)
			},
		},
	}
	var homeResults []EventResult
	var awayResults []EventResult
	var homeChances, awayChances, homeGoals, awayGoals int

	for i := 0; i < numberOfHomeEvents; i++ {
		event := homeEvents[rand.Intn(len(homeEvents))]
		log.Println("team event", event)
		result, newHomeChances, newAwayChances, newHomeGoals, newAwayGoals, err := event.Execute()
		if err != nil {
			fmt.Printf("Error executing home event: %v\n", err)
			continue
		}
		if result == "" {
			fmt.Println("Generated empty event for home!")
		} else {
			fmt.Printf("Generated home event: %s\n", result)
		}
		homeChances += newHomeChances
		awayChances += newAwayChances
		homeGoals += newHomeGoals
		awayGoals += newAwayGoals

		minute := rand.Intn(90)
		homeResults = append(homeResults, EventResult{
			Event:     result + fmt.Sprintf(" for the team %s", home.Name),
			Minute:    minute,
			EventType: event.Name,
			TeamId:    home.Id,
			TeamName:  fmt.Sprintf(" %s", home.Name),
		})
		fmt.Printf("Generated event: %s at minute %d\n", result, minute)

	}
	for i := 0; i < numberOfAwayEvents; i++ {
		event := awayEvents[rand.Intn(len(awayEvents))]
		log.Println("away event", event)
		result, newAwayChances, newHomeChances, newAwayGoals, newHomeGoals, err := event.Execute()
		if err != nil {
			fmt.Printf("Error executing away event: %v\n", err)
			continue
		}

		homeChances += newHomeChances
		awayChances += newAwayChances
		homeGoals += newHomeGoals
		awayGoals += newAwayGoals

		minute := rand.Intn(90)
		awayResults = append(awayResults, EventResult{
			Event:     result + " para " + awayHome.Name,
			Minute:    minute,
			EventType: event.Name,
			TeamId:    awayHome.Id,
			TeamName:  awayHome.Name,
		})
		fmt.Printf("Generated event: %s at minute %d\n", result, minute)

	}

	return MatchEventStats{
		HomeEvents:       homeResults,
		AwayEvents:       awayResults,
		HomeScoreChances: homeChances,
		AwayScoreChances: awayChances,
		HomeGoals:        homeGoals,
		AwayGoals:        awayGoals,
	}
}

func CalculateTotalQuality(homeTotalTechnique, homeTotalMental, homeTotalPhysique, awayTotalTechnique, awayTotalMental, awayTotalPhysique int) (int, int, int, error) {

	homeTotalQuality := homeTotalTechnique + homeTotalMental + homeTotalPhysique
	awayTotalQuality := awayTotalTechnique + awayTotalMental + awayTotalPhysique
	allQuality := homeTotalQuality + awayTotalQuality

	if allQuality == 0 {
		return 0, 0, 0, errors.New("error. quality cant be nil")
	}

	return homeTotalQuality, awayTotalQuality, allQuality, nil

}

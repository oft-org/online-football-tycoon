package match

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/robertobouses/online-football-tycoon/team"
)

func CalculateNumberOfMatchEvents(homeGameTempo, awayGameTempo string) (int, error) {

	var tempoMap = map[string]int{
		"slow_tempo":     1,
		"balanced_tempo": 2,
		"fast_tempo":     3,
	}

	lineupTempo := tempoMap[homeGameTempo]
	rivalTempo := tempoMap[awayGameTempo]
	matchTempo := lineupTempo + rivalTempo
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

func DistributeMatchEvents(lineup, rivalLineup team.Team, numberOfMatchEvents int) (int, int, error) {
	log.Println("lineup en DistributeMatchEvents", lineup)
	log.Println("rival en DistributeMatchEvents", rivalLineup)
	lineupTotalQuality, err := CalculateQuality(lineup)
	if err != nil {
		return 0, 0, err
	}
	log.Println("total lineup Quality", lineupTotalQuality)
	rivalTotalQuality, err := CalculateQuality(rivalLineup)
	if err != nil {
		return 0, 0, err
	}
	log.Println("total rival Quality", rivalTotalQuality)
	allQuality := lineupTotalQuality + rivalTotalQuality
	var lineupEvents int
	lineupProportion := float64(lineupTotalQuality) / float64(allQuality)

	lineupEvents = int(lineupProportion*float64(numberOfMatchEvents)) + rand.Intn(4) + 2

	log.Printf("number of lineup events %v ANTES DE RANDOMFACTOR", lineupEvents)

	randomFactor := rand.Intn(11) - 5

	lineupEvents += randomFactor

	rivalEvents := numberOfMatchEvents - lineupEvents
	log.Printf("number of lineup events %v, rival events %v Despues DE RANDOMFACTOR", lineupEvents, rivalEvents)
	if lineupEvents <= 0 {
		lineupEvents = 0
	}
	if rivalEvents < 0 {
		rivalEvents = 0
	}
	log.Printf("number of lineup events %v, rival events %v", lineupEvents, rivalEvents)
	return lineupEvents, rivalEvents, nil
}

func CalculateQuality(lineup team.Team) (int, error) {
	var totalTechnique, totalMental, totalPhysique int
	for _, player := range lineup.Players {
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

func GetRandomDefender(lineup []team.Player) *team.Player {
	var defenders []team.Player
	for _, player := range lineup {
		if player.Position == "defender" {
			defenders = append(defenders, player)
		}

	}
	return GetRandomPlayer(defenders)
}

func GetRandomMidfielder(lineup []team.Player) *team.Player {
	var midfielders []team.Player
	for _, player := range lineup {
		if player.Position == "midfielder" {
			midfielders = append(midfielders, player)
		}

	}
	return GetRandomPlayer(midfielders)
}

func GetRandomForward(lineup []team.Player) *team.Player {
	var forwards []team.Player
	for _, player := range lineup {
		if player.Position == "forward" {
			forwards = append(forwards, player)
		}

	}
	return GetRandomPlayer(forwards)
}

func GetGoalkeeper(lineup []team.Player) *team.Player {
	var goalkeepers []team.Player
	for _, player := range lineup {
		if player.Position == "goalkeeper" {
			goalkeepers = append(goalkeepers, player)
		}

	}
	return GetRandomPlayer(goalkeepers)
}

func GetRandomPlayerExcludingGoalkeeper(lineup []team.Player) *team.Player {
	var playersExcludingGoalkeepers []team.Player
	for _, player := range lineup {
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
	Event  string `json:"event"`
	Minute int    `json:"minute"`
	Team   string `json:"team"`
}

func GenerateEvents(lineup, rivalLineup team.Team, numberOfLineupEvents, numberOfRivalEvents int) ([]EventResult, []EventResult, int, int, int, int) {

	lineupEvents := []Event{
		{
			"Pase clave",
			func() (string, int, int, int, int, error) {
				return KeyPass(lineup, rivalLineup)
			},
		},
		{
			"Remate a puerta",
			func() (string, int, int, int, int, error) {
				return Shot(lineup, rivalLineup, GetRandomForward(lineup.Players))
			},
		},
		{
			"Penalty",
			func() (string, int, int, int, int, error) {
				return PenaltyKick(lineup, rivalLineup)
			},
		},
		{
			"Tiro lejano",
			func() (string, int, int, int, int, error) {
				return LongShot(lineup, rivalLineup)
			},
		},
		{
			" Lanzamiento de Falta Indirecta",
			func() (string, int, int, int, int, error) {
				return IndirectFreeKick(lineup, rivalLineup)
			},
		},
		{
			"Regate",
			func() (string, int, int, int, int, error) {
				return Dribble(lineup, rivalLineup)
			},
		},
		{
			"Falta",
			func() (string, int, int, int, int, error) {
				return Foul(lineup, rivalLineup, nil)
			},
		},

		{
			"Gran Ocasión",
			func() (string, int, int, int, int, error) {
				return GreatScoringChance(lineup)
			},
		},
		{
			"Córner",
			func() (string, int, int, int, int, error) {
				return CornerKick(lineup, rivalLineup)
			},
		},
		{
			"Fuera de Juego",
			func() (string, int, int, int, int, error) {
				return Offside(lineup, rivalLineup)
			},
		},
		{
			"Cabezazo",
			func() (string, int, int, int, int, error) {
				return Headed(lineup, rivalLineup)
			},
		}, {
			"Contragolpe",
			func() (string, int, int, int, int, error) {
				return CounterAttack(lineup, rivalLineup)
			},
		},
	}

	rivalEvents := []Event{
		{
			"Pase clave",
			func() (string, int, int, int, int, error) {
				return KeyPass(rivalLineup, lineup)
			},
		},
		{
			"Remate a puerta",
			func() (string, int, int, int, int, error) {
				return Shot(rivalLineup, lineup, GetRandomForward(rivalLineup.Players))
			},
		},
		{
			"Penalty",
			func() (string, int, int, int, int, error) {
				return PenaltyKick(rivalLineup, lineup)
			},
		},
		{
			"Tiro lejano",
			func() (string, int, int, int, int, error) {
				return LongShot(rivalLineup, lineup)
			},
		},
		{
			" Lanzamiento de Falta Indirecta",
			func() (string, int, int, int, int, error) {
				return IndirectFreeKick(rivalLineup, lineup)
			},
		},
		{
			"Regate",
			func() (string, int, int, int, int, error) {
				return Dribble(rivalLineup, lineup)
			},
		},
		{
			"Falta",
			func() (string, int, int, int, int, error) {
				return Foul(rivalLineup, lineup, nil)
			},
		},
		{
			"Gran Ocasión",
			func() (string, int, int, int, int, error) {
				return GreatScoringChance(rivalLineup)
			},
		},
		{
			"Córner",
			func() (string, int, int, int, int, error) {
				return CornerKick(rivalLineup, lineup)
			},
		},
		{
			"Fuera de Juego",
			func() (string, int, int, int, int, error) {
				return Offside(rivalLineup, lineup)
			},
		},
		{
			"Cabezazo",
			func() (string, int, int, int, int, error) {
				return Headed(rivalLineup, lineup)
			},
		}, {
			"Contragolpe",
			func() (string, int, int, int, int, error) {
				return CounterAttack(rivalLineup, lineup)
			},
		},
	}
	var lineupResults []EventResult
	var rivalResults []EventResult
	var lineupChances, rivalChances, lineupGoals, rivalGoals int

	for i := 0; i < numberOfLineupEvents; i++ {
		event := lineupEvents[rand.Intn(len(lineupEvents))]
		log.Println("evento de tu equipo", event)
		result, newLineupChances, newRivalChances, newLineupGoals, newRivalGoals, err := event.Execute()
		if err != nil {
			fmt.Printf("Error executing lineup event: %v\n", err)
			continue
		}
		if result == "" {
			fmt.Println("Generated empty event for lineup!")
		} else {
			fmt.Printf("Generated lineup event: %s\n", result)
		}
		lineupChances += newLineupChances
		rivalChances += newRivalChances
		lineupGoals += newLineupGoals
		rivalGoals += newRivalGoals

		minute := rand.Intn(90)
		lineupResults = append(lineupResults, EventResult{
			Event:  result + " para tu equipo",
			Minute: minute,
			Team:   " tu equipo",
		})
		fmt.Printf("Generated event: %s at minute %d\n", result, minute)

	}
	for i := 0; i < numberOfRivalEvents; i++ {
		event := rivalEvents[rand.Intn(len(rivalEvents))]
		log.Println("evento del rival", event)
		result, newRivalChances, newLineupChances, newRivalGoals, newLineupGoals, err := event.Execute()
		if err != nil {
			fmt.Printf("Error executing rival event: %v\n", err)
			continue
		}

		lineupChances += newLineupChances
		rivalChances += newRivalChances
		lineupGoals += newLineupGoals
		rivalGoals += newRivalGoals

		minute := rand.Intn(90)
		rivalResults = append(rivalResults, EventResult{
			Event:  result + " para " + rivalLineup.Name,
			Minute: minute,
			Team:   rivalLineup.Name,
		})
		fmt.Printf("Generated event: %s at minute %d\n", result, minute)

	}

	return lineupResults, rivalResults, lineupChances, rivalChances, lineupGoals, rivalGoals
}

func CalculateTotalQuality(lineupTotalTechnique, lineupTotalMental, lineupTotalPhysique, rivalTotalTechnique, rivalTotalMental, rivalTotalPhysique int) (int, int, int, error) {

	lineupTotalQuality := lineupTotalTechnique + lineupTotalMental + lineupTotalPhysique
	rivalTotalQuality := rivalTotalTechnique + rivalTotalMental + rivalTotalPhysique
	allQuality := lineupTotalQuality + rivalTotalQuality

	if allQuality == 0 {
		return 0, 0, 0, errors.New("error. quality cant be nil")
	}

	return lineupTotalQuality, rivalTotalQuality, allQuality, nil

}

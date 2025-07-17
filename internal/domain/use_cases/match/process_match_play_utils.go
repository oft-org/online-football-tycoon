package match

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/robertobouses/online-football-tycoon/internal/domain"
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

func DistributeMatchEvents(home, away domain.Team, numberOfMatchEvents int, homeFactorNumberEvents, awayFactorNumberEvents float64) (int, int, error) {
	const (
		homeEventMaxBonus       = 3
		homeEventBaseBonus      = 1
		homeEventRandomRange    = 5
		homeEventRandomOffset   = 3
		homeEventOverflowAdjust = 2
	)

	log.Println("home team in DistributeMatchEvents", home)
	log.Println("away team in DistributeMatchEvents", away)

	homeTotalQuality, err := CalculateQuality(home)
	if err != nil {
		return 0, 0, err
	}
	log.Println("total home Quality", homeTotalQuality)
	awayTotalQuality, err := CalculateQuality(away)
	if err != nil {
		return 0, 0, err
	}
	log.Println("total away Quality", awayTotalQuality)
	allQuality := homeTotalQuality + awayTotalQuality

	var homeEvents int
	homeProportion := float64(homeTotalQuality) / float64(allQuality)

	homeEvents = int(homeProportion*float64(numberOfMatchEvents)) + rand.Intn(homeEventMaxBonus) + homeEventBaseBonus

	log.Printf("number of home events %v BEFORE RANDOMFACTOR", homeEvents)

	homeEvents = homeEvents * int(homeFactorNumberEvents) / int(awayFactorNumberEvents)

	randomFactor := rand.Intn(homeEventRandomRange) - homeEventRandomOffset

	homeEvents += randomFactor

	if homeEvents > numberOfMatchEvents {
		homeEvents = numberOfMatchEvents - rand.Intn(homeEventOverflowAdjust)
	}

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

func CalculateQuality(home domain.Team) (int, error) {
	var totalTechnique, totalMental, totalPhysique int
	for _, player := range home.Players {
		totalTechnique += player.Technique
		totalMental += player.Mental
		totalPhysique += player.Physique
	}

	return 2*totalTechnique + 3*totalMental + 2*totalPhysique, nil
}

func GetRandomDefender(home []domain.Player) *domain.Player {
	var defenders []domain.Player
	for _, player := range home {
		if player.Position == domain.PositionDefender {
			defenders = append(defenders, player)
		}

	}
	return GetRandomPlayer(defenders)
}

func GetRandomMidfielder(home []domain.Player) *domain.Player {
	var midfielders []domain.Player
	for _, player := range home {
		if player.Position == domain.PositionMidfielder {
			midfielders = append(midfielders, player)
		}

	}
	return GetRandomPlayer(midfielders)
}

func GetRandomForward(home []domain.Player) *domain.Player {
	var forwards []domain.Player
	for _, player := range home {
		if player.Position == domain.PositionForward {
			forwards = append(forwards, player)
		}

	}
	return GetRandomPlayer(forwards)
}

func GetGoalkeeper(home []domain.Player) *domain.Player {
	var goalkeepers []domain.Player
	for _, player := range home {
		if player.Position == domain.PositionGoalkeeper {
			goalkeepers = append(goalkeepers, player)
		}

	}
	return GetRandomPlayer(goalkeepers)
}

func GetRandomPlayerExcludingGoalkeeper(home []domain.Player) *domain.Player {
	var playersExcludingGoalkeepers []domain.Player
	for _, player := range home {
		if player.Position != domain.PositionGoalkeeper {
			playersExcludingGoalkeepers = append(playersExcludingGoalkeepers, player)
		}

	}
	return GetRandomPlayer(playersExcludingGoalkeepers)
}

func GetRandomPlayer(filteredPlayers []domain.Player) *domain.Player {
	if len(filteredPlayers) == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	randomPlayer := filteredPlayers[rand.Intn(len(filteredPlayers))]
	return &randomPlayer
}

func GenerateEvents(home, awayHome domain.Team, numberOfHomeEvents, numberOfAwayEvents int) domain.MatchEventStats {

	homeEvents := []domain.Event{
		{
			string(EventTypeKeyPass),
			func() (string, int, int, int, int, error) {
				return KeyPass(home, awayHome)
			},
		},
		{
			string(EventTypeShot),
			func() (string, int, int, int, int, error) {
				return Shot(home, awayHome, GetRandomForward(home.Players))
			},
		},
		{
			string(EventTypePenaltyKick),
			func() (string, int, int, int, int, error) {
				return PenaltyKick(home, awayHome)
			},
		},
		{
			string(EventTypeLongShot),
			func() (string, int, int, int, int, error) {
				return LongShot(home, awayHome)
			},
		},
		{
			string(EventTypeIndirectFreeKick),
			func() (string, int, int, int, int, error) {
				return IndirectFreeKick(home, awayHome)
			},
		},
		{
			string(EventTypeDribble),
			func() (string, int, int, int, int, error) {
				return Dribble(home, awayHome)
			},
		},
		{
			string(EventTypeFoul),
			func() (string, int, int, int, int, error) {
				return Foul(home, awayHome, nil)
			},
		},

		{
			string(EventTypeGreatScoringChance),
			func() (string, int, int, int, int, error) {
				return GreatScoringChance(home)
			},
		},
		{
			string(EventTypeCornerKick),
			func() (string, int, int, int, int, error) {
				return CornerKick(home, awayHome)
			},
		},
		{
			string(EventTypeOffside),
			func() (string, int, int, int, int, error) {
				return Offside(home, awayHome)
			},
		},
		{
			string(EventTypeHeaded),
			func() (string, int, int, int, int, error) {
				return Headed(home, awayHome)
			},
		}, {
			string(EventTypeCounterAttack),
			func() (string, int, int, int, int, error) {
				return CounterAttack(home, awayHome)
			},
		},
	}

	awayEvents := []domain.Event{
		{
			string(EventTypeKeyPass),
			func() (string, int, int, int, int, error) {
				return KeyPass(awayHome, home)
			},
		},
		{
			string(EventTypeShot),
			func() (string, int, int, int, int, error) {
				return Shot(awayHome, home, GetRandomForward(awayHome.Players))
			},
		},
		{
			string(EventTypePenaltyKick),
			func() (string, int, int, int, int, error) {
				return PenaltyKick(awayHome, home)
			},
		},
		{
			string(EventTypeLongShot),
			func() (string, int, int, int, int, error) {
				return LongShot(awayHome, home)
			},
		},
		{
			string(EventTypeIndirectFreeKick),
			func() (string, int, int, int, int, error) {
				return IndirectFreeKick(awayHome, home)
			},
		},
		{
			string(EventTypeDribble),
			func() (string, int, int, int, int, error) {
				return Dribble(awayHome, home)
			},
		},
		{
			string(EventTypeFoul),
			func() (string, int, int, int, int, error) {
				return Foul(awayHome, home, nil)
			},
		},
		{
			string(EventTypeGreatScoringChance),
			func() (string, int, int, int, int, error) {
				return GreatScoringChance(awayHome)
			},
		},
		{
			string(EventTypeCornerKick),
			func() (string, int, int, int, int, error) {
				return CornerKick(awayHome, home)
			},
		},
		{
			string(EventTypeOffside),
			func() (string, int, int, int, int, error) {
				return Offside(awayHome, home)
			},
		},
		{
			string(EventTypeHeaded),
			func() (string, int, int, int, int, error) {
				return Headed(awayHome, home)
			},
		}, {
			string(EventTypeCounterAttack),
			func() (string, int, int, int, int, error) {
				return CounterAttack(awayHome, home)
			},
		},
	}
	var homeResults []domain.EventResult
	var awayResults []domain.EventResult
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
		homeResults = append(homeResults, domain.EventResult{
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
		awayResults = append(awayResults, domain.EventResult{
			Event:     result + " para " + awayHome.Name,
			Minute:    minute,
			EventType: event.Name,
			TeamId:    awayHome.Id,
			TeamName:  awayHome.Name,
		})
		fmt.Printf("Generated event: %s at minute %d\n", result, minute)

	}

	return domain.MatchEventStats{
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

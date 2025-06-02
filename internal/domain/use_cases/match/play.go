package match

import (
	"fmt"
	"log"
	"sort"

	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type Simulator struct{}

func NewSimulator() Simulator {
	return Simulator{}
}
func (s Simulator) Play(m *domain.Match) (domain.Result, []domain.EventResult, error) {
	homeLineup := m.HomeMatchStrategy.StrategyTeam.Players

	for count, player := range homeLineup {
		log.Printf("home lineup player #%d: %+v", count, player)
	}
	awayLineup := m.AwayMatchStrategy.StrategyTeam.Players
	for count, player := range awayLineup {
		log.Printf("Away lineup player #%d: %+v", count, player)
	}
	log.Printf("Home Strategy Team details: %+v", m.HomeMatchStrategy.StrategyTeam)
	log.Printf("Home Strategy Team Players: %+v", m.HomeMatchStrategy.StrategyTeam.Players)

	homeTeam := m.HomeMatchStrategy.StrategyTeam

	awayTeam := m.AwayMatchStrategy.StrategyTeam

	log.Printf("Rival Lineup (Team %s): %+v", awayTeam.Id, awayLineup)

	numberOfMatchEvents, err := CalculateNumberOfMatchEvents(m.HomeMatchStrategy.GameTempo, m.AwayMatchStrategy.GameTempo)
	if err != nil {
		log.Println("error on numberOfMatchEvents", err)
		return domain.Result{}, []domain.EventResult{}, err
	}
	log.Println("numberOfMatchEvents", numberOfMatchEvents)

	numberOfLineupEvents, numberOfRivalEvents, err := DistributeMatchEvents(m.HomeMatchStrategy.StrategyTeam, m.AwayMatchStrategy.StrategyTeam, numberOfMatchEvents)
	if err != nil {
		log.Println("error al distribuir numberOfMatchEvents", err)
		return domain.Result{}, []domain.EventResult{}, err
	}
	log.Println("numberOfLineupEvents, numberOfRivalEvents", numberOfLineupEvents, numberOfRivalEvents)

	matchEventStats := GenerateEvents(homeTeam, awayTeam, numberOfLineupEvents, numberOfRivalEvents)
	breakMatch := domain.EventResult{
		Minute:    45,
		EventType: string(EventTypeMatchBreak),
		Event:     "Descanso",
		TeamId:    homeTeam.Id,
	}

	endMatch := domain.EventResult{
		Minute:    90,
		EventType: string(EventTypeEndOfTheMatch),
		Event:     "Final del Partido",
		TeamId:    homeTeam.Id,
	}
	allEvents := append(matchEventStats.HomeEvents, matchEventStats.AwayEvents...)
	allEvents = append(allEvents, breakMatch, endMatch)
	sort.Slice(allEvents, func(i, j int) bool {
		return allEvents[i].Minute < allEvents[j].Minute
	})

	totalHomeTechnique, totalHomeMental, totalHomePhysique := totalStats(homeLineup)
	totalAwayTechnique, totalAwayMental, totalAwayPhysique := totalStats(awayLineup)

	strategy := m.HomeMatchStrategy

	resultOfStrategy, err := CalculateResultOfStrategy(homeLineup, strategy.Formation, strategy.PlayingStyle, strategy.GameTempo, strategy.PassingStyle, strategy.DefensivePositioning, strategy.BuildUpPlay, strategy.AttackFocus, strategy.KeyPlayerUsage)
	if err != nil {

		return domain.Result{}, []domain.EventResult{}, fmt.Errorf("error in calculating the result of the strategy: %w", err)
	}

	totalHomePhysique = totalHomePhysique + int(resultOfStrategy["teamPhysique"])

	lineupTotalQuality, rivalTotalQuality, allQuality, err := CalculateTotalQuality(totalHomeTechnique, totalHomeMental, totalHomePhysique, totalAwayTechnique, totalAwayMental, totalAwayPhysique)
	if err != nil {
		log.Println("Error calculating total quality:", err)
		return domain.Result{}, []domain.EventResult{}, err
	}
	log.Printf("Total Quality: player %d, rival %d, total quality %d\n", lineupTotalQuality, rivalTotalQuality, allQuality)
	lineupPercentagePossession, rivalPercentagePossession, err := CalculateBallPossession(totalHomeTechnique, totalHomeMental, lineupTotalQuality, rivalTotalQuality, allQuality, resultOfStrategy["teamPossession"])
	if err != nil {
		log.Println("Error CalculateBallPossession:", err)
		return domain.Result{}, []domain.EventResult{}, err
	}

	result := domain.Result{
		HomeStats: domain.TeamStats{
			BallPossession: lineupPercentagePossession,
			ScoringChances: matchEventStats.HomeScoreChances,
			Goals:          matchEventStats.HomeGoals,
		},
		AwayStats: domain.TeamStats{
			BallPossession: rivalPercentagePossession,
			ScoringChances: matchEventStats.AwayScoreChances,
			Goals:          matchEventStats.AwayGoals,
		},
	}

	return result, allEvents, nil
}

func totalStats(players []domain.Player) (technique, mental, physique int) {
	for _, p := range players {
		technique += p.Technique
		mental += p.Mental
		physique += p.Physique
	}
	return
}

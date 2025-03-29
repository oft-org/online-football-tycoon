package match

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/team"
)

type Match struct {
	HomeMatchStrategy Strategy
	AwayMatchStrategy Strategy
}

type Strategy struct {
	StrategyTeam         team.Team
	Formation            string
	PlayingStyle         string
	GameTempo            string
	PassingStyle         string
	DefensivePositioning string
	BuildUpPlay          string
	AttackFocus          string
	KeyPlayerUsage       string
}

type Result struct {
	HomeStats TeamStats
	AwayStats TeamStats
}

type TeamStats struct {
	BallPossession int
	ScoringChances int
	Goals          int
}

func (a AppService) PlayMatch(matchID uuid.UUID) (Result, error) {
	m, err := a.repo.GetMatchById(matchID)
	if err != nil {
		return Result{}, fmt.Errorf("error retrieving match: %w", err)
	}
	log.Println("______match info_______", m)

	result, err := m.Play()
	if err != nil {
		return Result{}, fmt.Errorf("error playing match: %w", err)
	}

	matchDate := time.Now()
	homeTeamId := m.HomeMatchStrategy.StrategyTeam.Id
	awayTeamId := m.AwayMatchStrategy.StrategyTeam.Id

	log.Println("homeTeamId, awayTeamId", homeTeamId, awayTeamId)

	err = a.repo.PostMatch(homeTeamId, awayTeamId, matchDate, result.HomeStats.Goals, result.AwayStats.Goals)
	if err != nil {
		return Result{}, fmt.Errorf("error PostMatch: %w", err)
	}

	return result, nil
}

func (m Match) Play() (Result, error) {
	lineup := m.HomeMatchStrategy.StrategyTeam.Players

	rivalLineup := m.AwayMatchStrategy.StrategyTeam.Players

	homeTeam := m.HomeMatchStrategy.StrategyTeam
	awayTeam := m.AwayMatchStrategy.StrategyTeam

	log.Println("rivalLineup", rivalLineup)

	numberOfMatchEvents, err := CalculateNumberOfMatchEvents(m.HomeMatchStrategy.GameTempo, m.AwayMatchStrategy.GameTempo)
	if err != nil {
		log.Println("error al obtener numberOfMatchEvents", err)
		return Result{}, err
	}
	log.Println("numberOfMatchEvents", numberOfMatchEvents)

	numberOfLineupEvents, numberOfRivalEvents, err := DistributeMatchEvents(m.HomeMatchStrategy.StrategyTeam, m.AwayMatchStrategy.StrategyTeam, numberOfMatchEvents)
	if err != nil {
		log.Println("error al distribuir numberOfMatchEvents", err)
		return Result{}, err
	}
	log.Println("numberOfLineupEvents, numberOfRivalEvents", numberOfLineupEvents, numberOfRivalEvents)

	homeEvents, awayEvents, homeScoreChances, awayScoreChances, homeGoals, awayGoals := GenerateEvents(homeTeam, awayTeam, numberOfLineupEvents, numberOfRivalEvents)
	breakMatch := EventResult{Minute: 45, Event: "Descanso"}
	endMatch := EventResult{Minute: 90, Event: "Final del Partido"}
	allEvents := append(homeEvents, awayEvents...)
	allEvents = append(allEvents, breakMatch, endMatch)
	sort.Slice(allEvents, func(i, j int) bool {
		return allEvents[i].Minute < allEvents[j].Minute
	})

	var totalHomeTechnique, totalHomeMental, totalHomePhysique int
	for _, player := range lineup {
		totalHomeTechnique += totalHomeTechnique + player.Technique
		totalHomeMental += totalHomeMental + player.Mental
		totalHomePhysique += totalHomePhysique + player.Physique
	}

	var totalAwayTechnique, totalAwayMental, totalAwayPhysique int
	for _, player := range rivalLineup {
		totalAwayTechnique += totalAwayTechnique + player.Technique
		totalAwayMental += totalAwayMental + player.Mental
		totalAwayPhysique += totalAwayPhysique + player.Physique
	}

	strategy := m.HomeMatchStrategy

	resultOfStrategy, err := CalculateResultOfStrategy(lineup, strategy.Formation, strategy.PlayingStyle, strategy.GameTempo, strategy.PassingStyle, strategy.DefensivePositioning, strategy.BuildUpPlay, strategy.AttackFocus, strategy.KeyPlayerUsage)
	if err != nil {
		log.Println("Error al calcular el resultado de la estrategia:", err)
		return Result{}, fmt.Errorf("error al calcular el resultado de la estrategia: %w", err)
	}

	totalHomePhysique = totalHomePhysique + int(resultOfStrategy["teamPhysique"])

	lineupTotalQuality, rivalTotalQuality, allQuality, err := CalculateTotalQuality(totalHomeTechnique, totalHomeMental, totalHomePhysique, totalAwayTechnique, totalAwayMental, totalAwayPhysique)
	if err != nil {
		log.Println("Error al calcular la calidad total:", err)
		return Result{}, err
	}
	log.Printf("Calidad total: jugador %d, rival %d, calidad total %d\n", lineupTotalQuality, rivalTotalQuality, allQuality)
	lineupPercentagePossession, rivalPercentagePossession, err := CalculateBallPossession(totalHomeTechnique, totalHomeMental, lineupTotalQuality, rivalTotalQuality, allQuality, resultOfStrategy["teamPossession"])
	if err != nil {
		log.Println("Error CalculateBallPossession:", err)
		return Result{}, err
	}

	result := Result{
		HomeStats: TeamStats{
			BallPossession: lineupPercentagePossession,
			ScoringChances: homeScoreChances,
			Goals:          homeGoals,
		},
		AwayStats: TeamStats{
			BallPossession: rivalPercentagePossession,
			ScoringChances: awayScoreChances,
			Goals:          awayGoals,
		},
	}
	return result, nil
}

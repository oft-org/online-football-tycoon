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

type MatchEventStats struct {
	HomeEvents       []EventResult
	AwayEvents       []EventResult
	HomeScoreChances int
	AwayScoreChances int
	HomeGoals        int
	AwayGoals        int
}

type MatchEventInfo struct {
	MatchID     uuid.UUID
	TeamId      uuid.UUID
	EventType   string
	Minute      int
	Description string
}

func (a AppService) PlayMatch(matchID uuid.UUID) (Result, error) {
	m, err := a.repo.GetMatchById(matchID)
	if err != nil {
		return Result{}, fmt.Errorf("error retrieving match: %w", err)
	}
	if m == nil {
		log.Printf("repo.GetMatchById returned nil for matchID: %s", matchID)
		return Result{}, fmt.Errorf("no match found with ID: %s", matchID)
	}
	result, allEvents, err := m.Play()
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

	for _, event := range allEvents {
		teamId, err := uuid.Parse(event.TeamId.String())
		if err != nil {
			log.Printf("Error parsing TeamId '%s': %v", event.TeamId, err)
			return Result{}, fmt.Errorf("invalid team ID '%s': %w", event.TeamId, err)
		}

		matchEventInfo := MatchEventInfo{
			MatchID:     matchID,
			TeamId:      teamId,
			EventType:   event.EventType,
			Minute:      event.Minute,
			Description: event.Event,
		}

		err = a.repo.PostMatchEvent(matchEventInfo)
		if err != nil {
			log.Printf("error posting event to repo: %v", err)
			return Result{}, fmt.Errorf("PostMatchEvent failed: %w", err)
		}
	}
	return result, nil
}

func (m Match) Play() (Result, []EventResult, error) {
	lineup := m.HomeMatchStrategy.StrategyTeam.Players

	for count, _ := range lineup {
		log.Println(count)
	}
	rivalLineup := m.AwayMatchStrategy.StrategyTeam.Players
	for count, _ := range rivalLineup {
		log.Println(count)
	}
	log.Printf("Buggg StrategyTeam: %+v", m.HomeMatchStrategy.StrategyTeam)
	log.Printf("Buggg Players: %+v", m.HomeMatchStrategy.StrategyTeam.Players)

	homeTeam := m.HomeMatchStrategy.StrategyTeam

	awayTeam := m.AwayMatchStrategy.StrategyTeam

	log.Println("rivalLineup", rivalLineup)

	numberOfMatchEvents, err := CalculateNumberOfMatchEvents(m.HomeMatchStrategy.GameTempo, m.AwayMatchStrategy.GameTempo)
	if err != nil {
		log.Println("error on numberOfMatchEvents", err)
		return Result{}, []EventResult{}, err
	}
	log.Println("numberOfMatchEvents", numberOfMatchEvents)

	numberOfLineupEvents, numberOfRivalEvents, err := DistributeMatchEvents(m.HomeMatchStrategy.StrategyTeam, m.AwayMatchStrategy.StrategyTeam, numberOfMatchEvents)
	if err != nil {
		log.Println("error al distribuir numberOfMatchEvents", err)
		return Result{}, []EventResult{}, err
	}
	log.Println("numberOfLineupEvents, numberOfRivalEvents", numberOfLineupEvents, numberOfRivalEvents)

	matchEventStats := GenerateEvents(homeTeam, awayTeam, numberOfLineupEvents, numberOfRivalEvents)
	breakMatch := EventResult{
		Minute:    45,
		EventType: string(EventTypeMatchBreak),
		Event:     "Descanso",
		TeamId:    homeTeam.Id,
	}

	endMatch := EventResult{
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

	var totalHomeTechnique, totalHomeMental, totalHomePhysique int
	for _, player := range lineup {
		totalHomeTechnique += player.Technique
		totalHomeMental += player.Mental
		totalHomePhysique += player.Physique
	}

	var totalAwayTechnique, totalAwayMental, totalAwayPhysique int
	for _, player := range rivalLineup {
		totalAwayTechnique += player.Technique
		totalAwayMental += player.Mental
		totalAwayPhysique += player.Physique
	}

	strategy := m.HomeMatchStrategy

	resultOfStrategy, err := CalculateResultOfStrategy(lineup, strategy.Formation, strategy.PlayingStyle, strategy.GameTempo, strategy.PassingStyle, strategy.DefensivePositioning, strategy.BuildUpPlay, strategy.AttackFocus, strategy.KeyPlayerUsage)
	if err != nil {

		return Result{}, []EventResult{}, fmt.Errorf("error in calculating the result of the strategy: %w", err)
	}

	totalHomePhysique = totalHomePhysique + int(resultOfStrategy["teamPhysique"])

	lineupTotalQuality, rivalTotalQuality, allQuality, err := CalculateTotalQuality(totalHomeTechnique, totalHomeMental, totalHomePhysique, totalAwayTechnique, totalAwayMental, totalAwayPhysique)
	if err != nil {
		log.Println("Error calculating total quality:", err)
		return Result{}, []EventResult{}, err
	}
	log.Printf("Total Quality: player %d, rival %d, total quality %d\n", lineupTotalQuality, rivalTotalQuality, allQuality)
	lineupPercentagePossession, rivalPercentagePossession, err := CalculateBallPossession(totalHomeTechnique, totalHomeMental, lineupTotalQuality, rivalTotalQuality, allQuality, resultOfStrategy["teamPossession"])
	if err != nil {
		log.Println("Error CalculateBallPossession:", err)
		return Result{}, []EventResult{}, err
	}

	result := Result{
		HomeStats: TeamStats{
			BallPossession: lineupPercentagePossession,
			ScoringChances: matchEventStats.HomeScoreChances,
			Goals:          matchEventStats.HomeGoals,
		},
		AwayStats: TeamStats{
			BallPossession: rivalPercentagePossession,
			ScoringChances: matchEventStats.AwayScoreChances,
			Goals:          matchEventStats.AwayGoals,
		},
	}

	return result, allEvents, nil
}

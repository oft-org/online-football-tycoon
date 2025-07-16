package match

import (
	"errors"
	"fmt"
	"sort"

	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

type strategyResult struct {
	homePossession float64
	homeChances    float64
	awayChances    float64
	homePhysique   int
}

type formationResult struct {
	homePossession float64
	homeChances    float64
	awayChances    float64
}

type playingStyleResult struct {
	homePossession float64
	homeChances    float64
	awayChances    float64
	physique       int
}

type gameTempoResult struct {
	homePossession float64
	homeChances    float64
	awayChances    float64
	physique       int
}

type passingStyleResult struct {
	homePossession float64
	awayChances    float64
}

type defensivePositioningResult struct {
	awayChances float64
	physique    int
}

type buildUpPlayResult struct {
	homePossession float64
}

type attackFocusResult struct {
	awayChances float64
}

type keyPlayerUsageResult struct {
	homePossession float64
	homeChances    float64
}

func CalculateResultOfStrategy(lineup []domain.Player, formation, playingStyle, gameTempo, passingStyle, defensivePositioning, buildUpPlay, attackFocus, keyPlayerUsage string) (result strategyResult, err error) {

	formationResult, err := CalculatePossessionChancesByFormation(lineup, formation)
	if err != nil {
		return strategyResult{}, fmt.Errorf("Error in formation: %v", err)
	}

	playingStyleResult, err := CalculatePossessionChancesByPlayingStyle(lineup, playingStyle)
	if err != nil {
		return strategyResult{}, fmt.Errorf("Error in playing style: %v", err)
	}

	gameTempoResult, err := CalculatePossessionChancesByGameTempo(gameTempo)
	if err != nil {
		return strategyResult{}, fmt.Errorf("Error in game tempo: %v", err)
	}

	passingStyleResult, err := CalculatePossessionChancesByPassingStyle(passingStyle)
	if err != nil {
		return strategyResult{}, fmt.Errorf("Error in passing style: %v", err)
	}

	defensivePositioningResult, err := CalculateRivalChancesByDefensivePositioning(lineup, defensivePositioning)
	if err != nil {
		return strategyResult{}, fmt.Errorf("Error in defensive positioning: %v", err)
	}

	buildUpPlayResult, err := CalculatePossessionByBuildUpPlay(lineup, buildUpPlay)
	if err != nil {
		return strategyResult{}, fmt.Errorf("Error in build-up play: %v", err)
	}

	attackFocusResult, err := CalculateRivalChancesByAttackFocus(lineup, attackFocus)
	if err != nil {
		return strategyResult{}, fmt.Errorf("Error in rival chances by attack focus: %v", err)
	}

	keyPlayerUsageResult, err := CalculateRivalChancesByKeyPlayerUsage(lineup, keyPlayerUsage)
	if err != nil {
		return strategyResult{}, fmt.Errorf("Error in rival chances by key player usage: %v", err)
	}

	totalHomePossession := (formationResult.homePossession + playingStyleResult.homePossession + gameTempoResult.homePossession + passingStyleResult.homePossession + buildUpPlayResult.homePossession + keyPlayerUsageResult.homePossession) / 6
	totalHomeChances := (formationResult.homeChances + playingStyleResult.homeChances + gameTempoResult.homeChances + keyPlayerUsageResult.homeChances) / 4
	totalAwayChances := (formationResult.awayChances + playingStyleResult.awayChances + gameTempoResult.awayChances + passingStyleResult.awayChances + defensivePositioningResult.awayChances + attackFocusResult.awayChances) / 6

	totalHomePhysique := playingStyleResult.physique + gameTempoResult.physique + defensivePositioningResult.physique

	result = strategyResult{
		homePossession: totalHomePossession,
		homeChances:    totalHomeChances,
		awayChances:    totalAwayChances,
		homePhysique:   totalHomePhysique,
	}
	return result, nil
}

func CalculatePossessionChancesByFormation(lineup []domain.Player, formation string) (result formationResult, err error) {

	totalDefendersQuality, err := getTwoBestPlayers(lineup, domain.PositionDefender)
	if err != nil {
		return formationResult{}, fmt.Errorf("Error getting two best defenders, in getTwoBestPlayers : %v", err)
	}
	totalMidfieldersQuality, err := getTwoBestPlayers(lineup, domain.PositionMidfielder)
	if err != nil {
		return formationResult{}, fmt.Errorf("Error getting two best midfielders, in getTwoBestPlayers: %v", err)
	}
	totalForwardersQuality, err := getTwoBestPlayers(lineup, domain.PositionForward)
	if err != nil {
		return formationResult{}, fmt.Errorf("Error getting two best forwarders, in getTwoBestPlayers: %v", err)
	}

	switch formation {
	case "4-4-2":
		if totalForwardersQuality >= 550 {
			result = formationResult{0.9, 1.2, 1}
		} else if totalForwardersQuality >= 360 {
			result = formationResult{0.9, 1, 1}
		} else {
			result = formationResult{0.8, 1, 1}
		}

	case "4-3-3":
		if totalMidfieldersQuality >= 510 {
			result = formationResult{0.9, 1.2, 1.1}
		} else {
			result = formationResult{0.8, 1.2, 1.1}
		}

	case "4-5-1":
		if totalMidfieldersQuality >= 540 {
			result = formationResult{1.4, 0.7, 0.7}
		} else if totalMidfieldersQuality >= 380 {
			result = formationResult{1.3, 0.7, 0.8}
		} else {
			result = formationResult{1.1, 0.6, 0.8}
		}

	case "5-4-1":
		if totalDefendersQuality >= 500 {
			result = formationResult{1, 0.5, 0.5}
		} else {
			result = formationResult{0.9, 0.5, 0.6}
		}

	case "5-3-2":
		if totalForwardersQuality >= 510 {
			result = formationResult{0.7, 1.1, 0.8}
		} else {
			result = formationResult{0.7, 1, 0.9}
		}

	case "3-4-3":
		if totalDefendersQuality >= 526 {
			result = formationResult{1.2, 1.3, 1.3}
		} else {
			result = formationResult{1, 1.3, 1.3}
		}

	case "3-5-2":
		if totalMidfieldersQuality >= 521 {
			result = formationResult{1.2, 1.1, 1.1}
		} else {
			result = formationResult{0.9, 1.1, 1.1}
		}

	default:
		return formationResult{}, errors.New("unknown formation")
	}

	return result, nil
}

func CalculatePossessionChancesByPlayingStyle(lineup []domain.Player, playingStyle string) (result playingStyleResult, err error) {
	totalDefendersQuality, err := getTwoBestPlayers(lineup, domain.PositionDefender)
	if err != nil {
		return playingStyleResult{}, fmt.Errorf("error getting two best defenders: %v", err)
	}

	totalMidfieldersQuality, err := getTwoBestPlayers(lineup, domain.PositionMidfielder)
	if err != nil {
		return playingStyleResult{}, fmt.Errorf("error getting two best midfielders: %v", err)
	}

	totalForwardersQuality, err := getTwoBestPlayers(lineup, domain.PositionForward)
	if err != nil {
		return playingStyleResult{}, fmt.Errorf("error getting two best forwarders: %v", err)
	}

	switch playingStyle {
	case "possession":
		if totalMidfieldersQuality >= 550 {
			result = playingStyleResult{1.6, 0.7, 0.8, 55}
		} else {
			result = playingStyleResult{1.4, 0.7, 0.8, 50}
		}
	case "counter_attack":
		if totalForwardersQuality >= 470 {
			result = playingStyleResult{0.7, 1.3, 0.9, -15}
		} else {
			result = playingStyleResult{0.7, 1.2, 0.9, -20}
		}

	case "direct_play":
		if totalForwardersQuality >= 400 {
			result = playingStyleResult{0.5, 1.1, 0.8, 20}
		} else {
			result = playingStyleResult{0.5, 1.0, 0.8, 10}
		}

	case "high_press":
		if totalMidfieldersQuality >= 440 {
			result = playingStyleResult{1.1, 1.4, 1.12, -190}
		} else {
			result = playingStyleResult{1.1, 1.35, 1.12, -220}
		}

	case "low_block":
		if totalDefendersQuality >= 410 {
			result = playingStyleResult{0.8, 0.4, 0.5, 130}
		} else {
			result = playingStyleResult{0.8, 0.3, 0.5, 110}
		}

	default:
		return playingStyleResult{}, errors.New("unknown playing style")
	}

	return result, nil
}

func CalculatePossessionChancesByGameTempo(gameTempo string) (result gameTempoResult, err error) {
	switch gameTempo {
	case "fast_tempo":
		result = gameTempoResult{0.8, 1.2, 1.1, -150}
	case "balanced_tempo":
		result = gameTempoResult{1, 1, 1, 10}
	case "slow_tempo":
		result = gameTempoResult{1.1, 0.6, 0.7, 250}

	default:
		return gameTempoResult{}, errors.New("unknown gameTempo")
	}
	return result, nil
}

func CalculatePossessionChancesByPassingStyle(passingStyle string) (result passingStyleResult, err error) {
	switch passingStyle {
	case "short":
		result = passingStyleResult{1.1, 1}
	case "long":
		result = passingStyleResult{0.8, 0.9}

	default:
		return passingStyleResult{}, errors.New("unknown passingStyle")
	}
	return result, nil
}

func CalculateRivalChancesByDefensivePositioning(lineup []domain.Player, defensivePositioning string) (result defensivePositioningResult, err error) {

	var totalMentalityOfDefenders, totalPhysiqueOfDefenders int

	for _, player := range lineup {
		if player.Position == domain.PositionDefender {
			totalMentalityOfDefenders += player.Mental
			totalPhysiqueOfDefenders += player.Physique
		}
	}

	switch defensivePositioning {
	case "zonal_marking":
		if totalMentalityOfDefenders >= 370 {
			result = defensivePositioningResult{0.7, 65}
		} else if totalMentalityOfDefenders >= 290 {
			result = defensivePositioningResult{0.9, 40}
		} else if totalMentalityOfDefenders >= 200 {
			result = defensivePositioningResult{1, 2}
		} else {
			result = defensivePositioningResult{1.45, -20}
		}
	case "man_marking":
		if totalMentalityOfDefenders >= 340 {
			result = defensivePositioningResult{0.8, 15}
		} else if totalMentalityOfDefenders >= 250 {
			result = defensivePositioningResult{0.9, 1}
		} else if totalPhysiqueOfDefenders >= 190 {
			result = defensivePositioningResult{1, -40}
		} else {
			result = defensivePositioningResult{1.3, -120}
		}
	default:
		return defensivePositioningResult{}, errors.New("unknown defensive positioning")
	}
	return result, nil
}

func CalculatePossessionByBuildUpPlay(lineup []domain.Player, buildUpPlay string) (result buildUpPlayResult, err error) {
	if len(lineup) == 0 {
		return buildUpPlayResult{}, errors.New("empty lineup")
	}
	var totalTechniqueOfGoalkeeper, totalMentalityOfGoalkeeper, totalTechniqueOfDefenders, totalMentalOfDefenders int
	var defenderCount int

	for _, player := range lineup {
		if player.Position == domain.PositionGoalkeeper {
			totalTechniqueOfGoalkeeper += player.Technique
			totalMentalityOfGoalkeeper += player.Mental
		} else if player.Position == domain.PositionDefender {
			totalTechniqueOfDefenders += player.Technique
			totalMentalOfDefenders += player.Mental
			defenderCount++

		}
	}

	if defenderCount == 0 {
		return buildUpPlayResult{}, errors.New("There are no defenders in the lineup")
	}

	totalQualityOfGoalkeeper := totalTechniqueOfGoalkeeper + totalMentalityOfGoalkeeper
	averageTotalQualityOfDefenders := (totalTechniqueOfDefenders + totalMentalOfDefenders) / defenderCount

	switch buildUpPlay {
	case "play_from_back":
		if totalTechniqueOfGoalkeeper >= 84 && totalMentalityOfGoalkeeper >= 84 && averageTotalQualityOfDefenders >= 79 {
			result = buildUpPlayResult{1.3}
		} else if totalTechniqueOfGoalkeeper >= 82 && totalMentalityOfGoalkeeper >= 82 || averageTotalQualityOfDefenders >= 70 && totalQualityOfGoalkeeper >= 150 {
			result = buildUpPlayResult{1.23}
		} else if totalQualityOfGoalkeeper >= 139 || averageTotalQualityOfDefenders >= 72 {
			result = buildUpPlayResult{1.10}
		} else if totalTechniqueOfGoalkeeper >= 66 || totalMentalityOfGoalkeeper >= 66 {
			result = buildUpPlayResult{1.07}
		} else {
			result = buildUpPlayResult{0.63}
		}

	case "long_clearance":

		if averageTotalQualityOfDefenders >= 86 {
			result = buildUpPlayResult{1.1}
		} else if averageTotalQualityOfDefenders >= 74 {
			result = buildUpPlayResult{1.02}
		}

		result = buildUpPlayResult{0.9}

	default:
		return buildUpPlayResult{}, errors.New("unknown buildUpPlay")
	}
	return result, nil
}

func CalculateRivalChancesByAttackFocus(lineup []domain.Player, attackFocus string) (result attackFocusResult, err error) {

	var totalTechniqueOfMidfield, totalPhysiqueOfMidfild int
	var forwardCount, midfieldersCount int

	for _, player := range lineup {
		if player.Position == domain.PositionMidfielder {
			totalTechniqueOfMidfield += player.Technique
			totalPhysiqueOfMidfild += player.Physique
			midfieldersCount++
		} else if player.Position == domain.PositionForward {
			forwardCount++

		}
	}

	if forwardCount == 0 {
		return attackFocusResult{}, errors.New("There are no forwarders in the lineup")
	}

	totalQualityOfMidfield := totalTechniqueOfMidfield + totalPhysiqueOfMidfild
	averageTotalQualityOfMidfield := totalQualityOfMidfield / midfieldersCount

	switch attackFocus {
	case "wide_play":
		if averageTotalQualityOfMidfield >= 84 && forwardCount >= 2 {
			result = attackFocusResult{1.28}
		} else if averageTotalQualityOfMidfield >= 82 {
			result = attackFocusResult{1.22}
		} else if totalQualityOfMidfield >= 245 || forwardCount >= 2 {
			result = attackFocusResult{1.09}
		} else if totalQualityOfMidfield >= 215 {
			result = attackFocusResult{1.06}
		} else {
			result = attackFocusResult{0.83}
		}

	case "central_play":
		if averageTotalQualityOfMidfield >= 79 && midfieldersCount >= 4 {
			result = attackFocusResult{1.21}
		} else if averageTotalQualityOfMidfield >= 76 {
			result = attackFocusResult{1.14}
		} else if midfieldersCount >= 4 {
			result = attackFocusResult{1.09}
		} else {

			result = attackFocusResult{0.91}
		}
	default:
		return attackFocusResult{}, errors.New("unknown AttackFocus")
	}
	return result, nil
}

func CalculateRivalChancesByKeyPlayerUsage(lineup []domain.Player, keyPlayerUsage string) (result keyPlayerUsageResult, err error) {

	var keyPlayer domain.Player

	for _, player := range lineup {
		if player.Technique+player.Mental+player.Physique > keyPlayer.Technique+keyPlayer.Mental+keyPlayer.Physique {
			keyPlayer = player
		}
	}

	totalQualityOfKeyPlayer := keyPlayer.Technique + keyPlayer.Mental + keyPlayer.Physique

	switch keyPlayerUsage {
	case "reference_player":
		if totalQualityOfKeyPlayer >= 278 {
			result = keyPlayerUsageResult{0.98, 1.9}
		} else if totalQualityOfKeyPlayer >= 271 {
			result = keyPlayerUsageResult{0.94, 1.64}
		} else if totalQualityOfKeyPlayer >= 254 {
			result = keyPlayerUsageResult{1, 1.52}
		} else if totalQualityOfKeyPlayer >= 216 {
			result = keyPlayerUsageResult{1, 1.26}
		} else if totalQualityOfKeyPlayer >= 204 {
			result = keyPlayerUsageResult{0.98, 1.1}
		} else {

			result = keyPlayerUsageResult{1.1, 0.67}
		}

	case "free_role_player":
		result = keyPlayerUsageResult{1.3, 0.98}

	default:
		return keyPlayerUsageResult{}, errors.New("unknown KeyPlayerUsage")
	}
	return result, nil
}

func getTwoBestPlayers(players []domain.Player, position string) (int, error) {
	var selected []domain.Player
	var midfielders []domain.Player

	for _, p := range players {
		if p.Position == position {
			selected = append(selected, p)
		} else if p.Position == domain.PositionMidfielder {
			midfielders = append(midfielders, p)
		}
	}

	if len(selected) < 2 {
		needed := 2 - len(selected)

		sort.Slice(midfielders, func(i, j int) bool {
			return playerTotalQuality(midfielders[i]) > playerTotalQuality(midfielders[j])
		})

		if len(midfielders) < needed {
			return 0, fmt.Errorf("not enough players in position %s and not enough midfielders to compensate", position)
		}
		selected = append(selected, midfielders[:needed]...)
	}

	sort.Slice(selected, func(i, j int) bool {
		return playerTotalQuality(selected[i]) > playerTotalQuality(selected[j])
	})

	sum := playerTotalQuality(selected[0]) + playerTotalQuality(selected[1])

	return sum, nil
}

func playerTotalQuality(p domain.Player) int {
	return p.Technique + p.Mental + p.Physique
}

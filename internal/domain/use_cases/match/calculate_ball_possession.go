package match

import (
	"log"
	"math/rand"
	"time"
)

func CalculateBallPossession(homeTotalTechnique, awayTotalTechnique, homeTotalQuality, awayTotalQuality, allQuality int, homePossessionResultOfStrategy, awayPossessionResultOfStrategy float64) (int, int, error) {
	percentageHomeQuality := (float64(homeTotalQuality) / float64(allQuality)) * 100

	switch {
	case float64(homeTotalTechnique) >= float64(awayTotalTechnique)*1.5:
		percentageHomeQuality *= 1.25
	case float64(homeTotalTechnique) >= float64(awayTotalTechnique)*1.4:
		percentageHomeQuality *= 1.2
	case float64(homeTotalTechnique) >= float64(awayTotalTechnique)*1.3:
		percentageHomeQuality *= 1.15
	case float64(homeTotalTechnique) >= float64(awayTotalTechnique)*1.2:
		percentageHomeQuality *= 1.1
	case float64(homeTotalTechnique) >= float64(awayTotalTechnique)*1.1:
		percentageHomeQuality *= 1.05
	}

	switch {
	case float64(homeTotalTechnique) <= float64(awayTotalTechnique)*1.5:
		percentageHomeQuality /= 1.25
	case float64(homeTotalTechnique) <= float64(awayTotalTechnique)*1.4:
		percentageHomeQuality /= 1.2
	case float64(homeTotalTechnique) <= float64(awayTotalTechnique)*1.3:
		percentageHomeQuality /= 1.15
	case float64(homeTotalTechnique) <= float64(awayTotalTechnique)*1.2:
		percentageHomeQuality /= 1.1
	case float64(homeTotalTechnique) <= float64(awayTotalTechnique)*1.1:
		percentageHomeQuality /= 1.05
	}
	log.Println("team possession before strategy", percentageHomeQuality)

	if homePossessionResultOfStrategy >= awayPossessionResultOfStrategy {
		percentageHomeQuality = percentageHomeQuality * homePossessionResultOfStrategy
	} else {
		percentageHomeQuality = percentageHomeQuality / awayPossessionResultOfStrategy
	}
	log.Println("team possession after strategy", percentageHomeQuality)

	rand.Seed(time.Now().UnixNano())
	randomFactor := 0.8 + rand.Float64()*(1.2-0.8)
	log.Println("randomFactor is", randomFactor)
	percentageHomeQualityWithRandomFactor := percentageHomeQuality * randomFactor
	log.Println("team possession after randomFactor", percentageHomeQualityWithRandomFactor)

	if percentageHomeQualityWithRandomFactor <= 45 {
		percentageHomeQualityWithRandomFactor *= 1.22
	} else if percentageHomeQualityWithRandomFactor <= 54 {
		percentageHomeQualityWithRandomFactor *= 1.15
	} else if percentageHomeQualityWithRandomFactor <= 67 {
		percentageHomeQualityWithRandomFactor *= 1.07
	}

	if percentageHomeQualityWithRandomFactor > 83 {
		percentageHomeQualityWithRandomFactor = 83
	} else if percentageHomeQualityWithRandomFactor < 17 {
		percentageHomeQualityWithRandomFactor = 17
	}

	percentageHome := int(percentageHomeQualityWithRandomFactor)
	percentageAway := 100 - percentageHome

	return percentageHome, percentageAway, nil
}

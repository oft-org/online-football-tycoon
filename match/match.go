package match

import (
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

func (m Match) Play() (Result, error) {
	var result Result
	var totalHomeGoals, totalAwayGoals int

	homeTotalTechnique, _, _, err := m.HomeMatchStrategy.StrategyTeam.CalculateTotalSkillsByTeam()
	if err != nil {
		return Result{}, err
	}

	awayTotalTechnique, _, _, err := m.AwayMatchStrategy.StrategyTeam.CalculateTotalSkillsByTeam()
	if err != nil {
		return Result{}, err
	}

	if homeTotalTechnique > awayTotalTechnique {
		totalHomeGoals += 1
	}
	if homeTotalTechnique < awayTotalTechnique {
		totalAwayGoals += 1
	} else {
		totalHomeGoals += 1
		totalAwayGoals += 1
	}

	result.HomeStats.Goals = totalHomeGoals
	result.AwayStats.Goals = totalAwayGoals

	return result, nil
}

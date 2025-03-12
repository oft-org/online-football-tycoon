package match

import "github.com/robertobouses/online-football-tycoon/team"

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
	if m.HomeMatchStrategy.GameTempo == "fast" {
		result.HomeStats.Goals = 2
		result.AwayStats.Goals = 0
	}
	return result, nil
}

package repository

import (
	"github.com/robertobouses/online-football-tycoon/match"
	"github.com/robertobouses/online-football-tycoon/team"
)

func (r *repository) GetMatches() ([]match.Match, error) {
	rows, err := r.getMatches.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []match.Match
	for rows.Next() {
		var m match.Match
		var homeTeam, awayTeam team.Team
		var homeStrategy, awayStrategy match.Strategy

		err := rows.Scan(
			&homeTeam.Name,
			&awayTeam.Name,
			&homeStrategy.Formation,
			&homeStrategy.PlayingStyle,
			&homeStrategy.GameTempo,
			&homeStrategy.PassingStyle,
			&homeStrategy.DefensivePositioning,
			&homeStrategy.BuildUpPlay,
			&homeStrategy.AttackFocus,
			&homeStrategy.KeyPlayerUsage,
			&awayStrategy.Formation,
			&awayStrategy.PlayingStyle,
			&awayStrategy.GameTempo,
			&awayStrategy.PassingStyle,
			&awayStrategy.DefensivePositioning,
			&awayStrategy.BuildUpPlay,
			&awayStrategy.AttackFocus,
			&awayStrategy.KeyPlayerUsage,
		)
		if err != nil {
			return nil, err
		}

		homeStrategy.StrategyTeam = homeTeam
		awayStrategy.StrategyTeam = awayTeam

		m.HomeMatchStrategy = homeStrategy
		m.AwayMatchStrategy = awayStrategy

		matches = append(matches, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

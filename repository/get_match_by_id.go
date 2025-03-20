package repository

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/match"
	"github.com/robertobouses/online-football-tycoon/team"
)

func (r *repository) GetMatchById(matchId uuid.UUID) (*match.Match, error) {
	row := r.getMatchById.QueryRow(matchId)

	var m match.Match
	var homeTeam, awayTeam team.Team
	var homeStrategy, awayStrategy match.Strategy
	var id uuid.UUID

	err := row.Scan(
		&id,
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

	return &m, nil
}

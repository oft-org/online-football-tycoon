package match

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) GetMatches() ([]domain.Match, error) {
	rows, err := r.getMatches.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []domain.Match
	for rows.Next() {
		var m domain.Match
		var homeTeam, awayTeam domain.Team
		var homeStrategy, awayStrategy domain.Strategy
		var matchId uuid.UUID

		err := rows.Scan(
			&matchId,
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

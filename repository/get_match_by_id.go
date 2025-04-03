package repository

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/match"
	"github.com/robertobouses/online-football-tycoon/team"
)

func (r *repository) GetMatchById(matchId uuid.UUID) (*match.Match, error) {
	rows, err := r.getMatchById.Query(matchId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var m match.Match
	var homeTeam, awayTeam team.Team
	var homeStrategy, awayStrategy match.Strategy
	var id, homeTeamId, awayTeamId uuid.UUID

	for rows.Next() {
		var homePlayer, awayPlayer team.Player
		err := rows.Scan(
			&id,
			&homeTeamId,
			&homeTeam.Name,
			&awayTeamId,
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

			&homePlayer.PlayerId,
			&homePlayer.FirstName,
			&homePlayer.LastName,
			&homePlayer.Position,
			&homePlayer.Technique,
			&homePlayer.Mental,
			&homePlayer.Physique,

			&awayPlayer.PlayerId,
			&awayPlayer.FirstName,
			&awayPlayer.LastName,
			&awayPlayer.Position,
			&awayPlayer.Technique,
			&awayPlayer.Mental,
			&awayPlayer.Physique,
		)
		if err != nil {
			return nil, err
		}

		homeTeam.Id = homeTeamId
		awayTeam.Id = awayTeamId

		homeTeam.Players = append(homeTeam.Players, homePlayer)
		awayTeam.Players = append(awayTeam.Players, awayPlayer)
	}
	homeStrategy.StrategyTeam = homeTeam
	awayStrategy.StrategyTeam = awayTeam

	m.HomeMatchStrategy = homeStrategy
	m.AwayMatchStrategy = awayStrategy

	return &m, nil
}

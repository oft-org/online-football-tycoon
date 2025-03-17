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
			&homeTeam.Name,                     // home_team_name
			&awayTeam.Name,                     // away_team_name
			&homeStrategy.Formation,            // home_formation
			&homeStrategy.PlayingStyle,         // home_playing_style
			&homeStrategy.GameTempo,            // home_game_tempo
			&homeStrategy.PassingStyle,         // home_passing_style
			&homeStrategy.DefensivePositioning, // home_defensive_positioning
			&homeStrategy.BuildUpPlay,          // home_build_up_play
			&homeStrategy.AttackFocus,          // home_attack_focus
			&homeStrategy.KeyPlayerUsage,       // home_key_player_usage
			&awayStrategy.Formation,            // away_formation
			&awayStrategy.PlayingStyle,         // away_playing_style
			&awayStrategy.GameTempo,            // away_game_tempo
			&awayStrategy.PassingStyle,         // away_passing_style
			&awayStrategy.DefensivePositioning, // away_defensive_positioning
			&awayStrategy.BuildUpPlay,          // away_build_up_play
			&awayStrategy.AttackFocus,          // away_attack_focus
			&awayStrategy.KeyPlayerUsage,       // away_key_player_usage

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

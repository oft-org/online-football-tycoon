package match

import (
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (r *Repository) GetMatchStrategyById(matchId uuid.UUID) (*domain.Match, error) {
	var m domain.Match
	var homeTeam, awayTeam domain.Team
	var homeStrategy, awayStrategy domain.Strategy

	row := r.getMatchTeams.QueryRow(matchId)
	if err := row.Scan(
		&homeTeam.Id,
		&homeTeam.Name,
		&awayTeam.Id,
		&awayTeam.Name,
	); err != nil {
		return nil, err
	}

	row = r.getMatchStrategies.QueryRow(homeTeam.Id)
	if err := row.Scan(
		&homeStrategy.Formation,
		&homeStrategy.PlayingStyle,
		&homeStrategy.GameTempo,
		&homeStrategy.PassingStyle,
		&homeStrategy.DefensivePositioning,
		&homeStrategy.BuildUpPlay,
		&homeStrategy.AttackFocus,
		&homeStrategy.KeyPlayerUsage,
	); err != nil {
		return nil, err
	}

	row = r.getMatchStrategies.QueryRow(awayTeam.Id)
	if err := row.Scan(
		&awayStrategy.Formation,
		&awayStrategy.PlayingStyle,
		&awayStrategy.GameTempo,
		&awayStrategy.PassingStyle,
		&awayStrategy.DefensivePositioning,
		&awayStrategy.BuildUpPlay,
		&awayStrategy.AttackFocus,
		&awayStrategy.KeyPlayerUsage,
	); err != nil {
		return nil, err
	}

	rows, err := r.getMatchPlayers.Query(homeTeam.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		log.Println("GetMatchStrategyById: iterating over row of players")
		var homePlayer domain.Player
		if err := rows.Scan(
			&homePlayer.PlayerId,
			&homePlayer.FirstName,
			&homePlayer.LastName,
			&homePlayer.Position,
			&homePlayer.Technique,
			&homePlayer.Mental,
			&homePlayer.Physique,
		); err != nil {
			log.Printf("GetMatchStrategyById: error scanning player: %v", err)
			return nil, err
		}
		log.Printf("Local Player: %+v", homePlayer)
		homeTeam.Players = append(homeTeam.Players, homePlayer)
	}

	rows, err = r.getMatchPlayers.Query(awayTeam.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var awayPlayer domain.Player
		if err := rows.Scan(
			&awayPlayer.PlayerId,
			&awayPlayer.FirstName,
			&awayPlayer.LastName,
			&awayPlayer.Position,
			&awayPlayer.Technique,
			&awayPlayer.Mental,
			&awayPlayer.Physique,
		); err != nil {
			log.Printf("GetMatchStrategyById: error scanning player: %v", err)
			return nil, err
		}
		log.Printf("Visiting Player: %+v", awayPlayer)
		awayTeam.Players = append(awayTeam.Players, awayPlayer)
	}

	homeStrategy.StrategyTeam = homeTeam
	awayStrategy.StrategyTeam = awayTeam
	log.Printf("Total players on local team: %d", len(homeTeam.Players))
	log.Printf("Total players on visiting team: %d", len(awayTeam.Players))

	m.HomeMatchStrategy = homeStrategy
	m.AwayMatchStrategy = awayStrategy

	return &m, nil
}

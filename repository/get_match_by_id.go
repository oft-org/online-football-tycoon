package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/match"
	"github.com/robertobouses/online-football-tycoon/team"
)

func (r *repository) GetMatchById(matchId uuid.UUID) (*match.Match, error) {
	var m match.Match
	var homeTeam, awayTeam team.Team
	var homeStrategy, awayStrategy match.Strategy

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
		log.Println("GetMatchById: iterando sobre fila de jugadores")
		var homePlayer team.Player
		if err := rows.Scan(
			&homePlayer.PlayerId,
			&homePlayer.FirstName,
			&homePlayer.LastName,
			&homePlayer.Position,
			&homePlayer.Technique,
			&homePlayer.Mental,
			&homePlayer.Physique,
		); err != nil {
			log.Printf("GetMatchById: error escaneando jugador: %v", err)
			return nil, err
		}
		log.Printf("Jugador local: %+v", homePlayer)
		homeTeam.Players = append(homeTeam.Players, homePlayer)
	}

	rows, err = r.getMatchPlayers.Query(awayTeam.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var awayPlayer team.Player
		if err := rows.Scan(
			&awayPlayer.PlayerId,
			&awayPlayer.FirstName,
			&awayPlayer.LastName,
			&awayPlayer.Position,
			&awayPlayer.Technique,
			&awayPlayer.Mental,
			&awayPlayer.Physique,
		); err != nil {
			log.Printf("GetMatchById: error escaneando jugador: %v", err)
			return nil, err
		}
		log.Printf("Jugador visitante: %+v", awayPlayer)
		awayTeam.Players = append(awayTeam.Players, awayPlayer)
	}

	homeStrategy.StrategyTeam = homeTeam
	awayStrategy.StrategyTeam = awayTeam
	log.Printf("Total jugadores en equipo local: %d", len(homeTeam.Players))
	log.Printf("Total jugadores en equipo visitante: %d", len(awayTeam.Players))

	m.HomeMatchStrategy = homeStrategy
	m.AwayMatchStrategy = awayStrategy

	return &m, nil
}

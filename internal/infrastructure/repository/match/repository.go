package match

import (
	"database/sql"
	"log"

	_ "embed"
)

//go:embed sql/get_matches.sql
var getMatchesQuery string

//go:embed sql/get_match_teams.sql
var getMatchTeamsQuery string

//go:embed sql/get_match_strategies.sql
var getMatchStrategiesQuery string

//go:embed sql/get_match_players.sql
var getMatchPlayersQuery string

//go:embed sql/post_match.sql
var postMatchQuery string

//go:embed sql/post_match_events.sql
var postMatchEventsQuery string

//go:embed sql/get_pending_matches.sql
var getPendingMatchesQuery string

//go:embed sql/get_match_by_id.sql
var getMatchByIDQuery string

//go:embed sql/update_match.sql
var updateMatchQuery string

//go:embed sql/get_match_events.sql
var getMatchEventsQuery string

//go:embed sql/get_season_matches.sql
var getSeasonMatchesQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	getMatchesStmt, err := db.Prepare(getMatchesQuery)
	if err != nil {
		return nil, err
	}

	getMatchTeamsStmt, err := db.Prepare(getMatchTeamsQuery)
	if err != nil {
		return nil, err
	}

	getMatchStrategiesStmt, err := db.Prepare(getMatchStrategiesQuery)
	if err != nil {
		return nil, err
	}

	getMatchPlayersStmt, err := db.Prepare(getMatchPlayersQuery)
	if err != nil {
		return nil, err
	}

	postMatchStmt, err := db.Prepare(postMatchQuery)
	if err != nil {
		return nil, err
	}

	postMatchEventsStmt, err := db.Prepare(postMatchEventsQuery)
	if err != nil {
		return nil, err
	}
	getPendingMatchesStmt, err := db.Prepare(getPendingMatchesQuery)
	if err != nil {
		return nil, err
	}

	getMatchByIDStmt, err := db.Prepare(getMatchByIDQuery)
	if err != nil {
		return nil, err
	}

	log.Println("Preparing updateMatch with SQL:")
	log.Println(updateMatchQuery)

	updateMatchStmt, err := db.Prepare(updateMatchQuery)
	if err != nil {
		return nil, err
	}
	getMatchEventsStmt, err := db.Prepare(getMatchEventsQuery)
	if err != nil {
		return nil, err
	}
	getSesaonMatchesStmt, err := db.Prepare(getSeasonMatchesQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                 db,
		getMatches:         getMatchesStmt,
		getMatchTeams:      getMatchTeamsStmt,
		getMatchStrategies: getMatchStrategiesStmt,
		getMatchPlayers:    getMatchPlayersStmt,
		postMatch:          postMatchStmt,
		postMatchEvents:    postMatchEventsStmt,
		getPendingMatches:  getPendingMatchesStmt,
		getMatchByID:       getMatchByIDStmt,
		updateMatch:        updateMatchStmt,
		getMatchEvents:     getMatchEventsStmt,
		getSeasonMatches:   getSesaonMatchesStmt,
	}, nil
}

type Repository struct {
	db                 *sql.DB
	getMatches         *sql.Stmt
	getMatchTeams      *sql.Stmt
	getMatchStrategies *sql.Stmt
	getMatchPlayers    *sql.Stmt
	postMatch          *sql.Stmt
	postMatchEvents    *sql.Stmt
	getPendingMatches  *sql.Stmt
	getMatchByID       *sql.Stmt
	updateMatch        *sql.Stmt
	getMatchEvents     *sql.Stmt
	getSeasonMatches   *sql.Stmt
}

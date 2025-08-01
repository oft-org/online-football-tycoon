package team

import (
	"math/rand/v2"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (a AppService) GenerateSeason(seasonID uuid.UUID, startDate time.Time) error {
	var matches []domain.SeasonMatch

	tournament, err := a.tournamentRepo.GetTournamentBySeasonID(seasonID)
	if err != nil {
		return err
	}
	teamIDs, err := a.repo.GetSeasonTeam(seasonID)
	if err != nil {
		return err
	}

	switch tournament.Type {
	case domain.TournamentLeague:
		matches = generateLeague(seasonID, teamIDs, startDate)

	case domain.TournamentCup:
		matches = generateCup(seasonID, teamIDs, startDate)
	}

	return a.matchRepo.PostMatches(matches)
}

func generateLeague(seasonID uuid.UUID, teamIDs []uuid.UUID, startDate time.Time) []domain.SeasonMatch {
	if len(teamIDs)%2 != 0 {
		teamIDs = append(teamIDs, uuid.Nil)
	}

	numRounds := len(teamIDs) - 1
	var matches []domain.SeasonMatch
	matchDay := 0

	for round := 0; round < numRounds; round++ {
		for i := 0; i < len(teamIDs)/2; i++ {
			home := teamIDs[i]
			away := teamIDs[len(teamIDs)-1-i]
			if home != uuid.Nil && away != uuid.Nil {
				matchDate := startDate.AddDate(0, 0, matchDay*7)

				matches = append(matches, domain.SeasonMatch{
					SeasonID:   seasonID,
					HomeTeamID: home,
					AwayTeamID: away,
					MatchDate:  matchDate,
				})
				matchDay++

				matchDate = startDate.AddDate(0, 0, matchDay*7)
				matches = append(matches, domain.SeasonMatch{
					SeasonID:   seasonID,
					HomeTeamID: away,
					AwayTeamID: home,
					MatchDate:  matchDate,
				})
				matchDay++
			}
		}
		teamIDs = append([]uuid.UUID{teamIDs[0]},
			append([]uuid.UUID{teamIDs[len(teamIDs)-1]}, teamIDs[1:len(teamIDs)-1]...)...,
		)
	}

	return matches
}

func generateCup(seasonID uuid.UUID, teamIDs []uuid.UUID, startDate time.Time) []domain.SeasonMatch {
	rand.Shuffle(len(teamIDs), func(i, j int) {
		teamIDs[i], teamIDs[j] = teamIDs[j], teamIDs[i]
	})

	var matches []domain.SeasonMatch
	for i := 0; i+1 < len(teamIDs); i += 2 {
		matchDate := startDate.AddDate(0, 0, (i/2)*7)
		match := domain.SeasonMatch{
			SeasonID:   seasonID,
			HomeTeamID: teamIDs[i],
			AwayTeamID: teamIDs[i+1],
			MatchDate:  matchDate,
		}
		matches = append(matches, match)
	}

	return matches
}

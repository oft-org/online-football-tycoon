package team

import (
	"math/rand/v2"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (a AppService) GenerateSeason(seasonID uuid.UUID) error {
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
		matches = generateLeague(seasonID, teamIDs)

	case domain.TournamentCup:
		matches = generateCup(seasonID, teamIDs)
	}

	if err := a.matchRepo.PostMatches(matches); err != nil {
		return err
	}
	return nil
}

func generateLeague(seasonID uuid.UUID, teamIDs []uuid.UUID) []domain.SeasonMatch {
	if len(teamIDs)%2 != 0 {
		teamIDs = append(teamIDs, uuid.Nil)
	}

	numRounds := len(teamIDs) - 1
	var matches []domain.SeasonMatch

	for round := 0; round < numRounds; round++ {
		for i := 0; i < len(teamIDs)/2; i++ {
			home := teamIDs[i]
			away := teamIDs[len(teamIDs)-1-i]
			if home != uuid.Nil && away != uuid.Nil {

				matches = append(matches, domain.SeasonMatch{
					SeasonID:   seasonID,
					HomeTeamID: home,
					AwayTeamID: away,
				})
				matches = append(matches, domain.SeasonMatch{
					SeasonID:   seasonID,
					HomeTeamID: away,
					AwayTeamID: home,
				})
			}
		}
		teamIDs = append([]uuid.UUID{teamIDs[0]},
			append([]uuid.UUID{teamIDs[len(teamIDs)-1]}, teamIDs[1:len(teamIDs)-1]...)...,
		)
	}

	return matches
}

func generateCup(seasonID uuid.UUID, teamIDs []uuid.UUID) []domain.SeasonMatch {
	rand.Shuffle(len(teamIDs), func(i, j int) {
		teamIDs[i], teamIDs[j] = teamIDs[j], teamIDs[i]
	})

	var matches []domain.SeasonMatch
	for i := 0; i+1 < len(teamIDs); i += 2 {
		match := domain.SeasonMatch{
			SeasonID:   seasonID,
			HomeTeamID: teamIDs[i],
			AwayTeamID: teamIDs[i+1],
		}
		matches = append(matches, match)
	}

	return matches
}

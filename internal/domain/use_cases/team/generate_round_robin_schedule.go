package team

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (a AppService) GenerateRoundRobinSchedule(seasonID uuid.UUID) error {
	teamIDs, err := a.repo.GetSeasonTeam(seasonID)
	if err != nil {
		return err
	}

	matches := generateMatchesForSeason(seasonID, teamIDs)

	if err := a.matchRepo.PostMatches(matches); err != nil {
		return err
	}
	return nil
}

func generateMatchesForSeason(seasonID uuid.UUID, teamIDs []uuid.UUID) []domain.SeasonMatch {
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
			}
		}
		teamIDs = append([]uuid.UUID{teamIDs[0]},
			append([]uuid.UUID{teamIDs[len(teamIDs)-1]}, teamIDs[1:len(teamIDs)-1]...)...)
	}

	for round := 0; round < numRounds; round++ {
		for i := 0; i < len(teamIDs)/2; i++ {
			home := teamIDs[len(teamIDs)-1-i]
			away := teamIDs[i]
			if home != uuid.Nil && away != uuid.Nil {
				matches = append(matches, domain.SeasonMatch{
					SeasonID:   seasonID,
					HomeTeamID: home,
					AwayTeamID: away,
				})
			}
		}

		teamIDs = append([]uuid.UUID{teamIDs[0]},
			append([]uuid.UUID{teamIDs[len(teamIDs)-1]}, teamIDs[1:len(teamIDs)-1]...)...)
	}

	return matches
}

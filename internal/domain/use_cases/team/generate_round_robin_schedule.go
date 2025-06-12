package team

import (
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (a AppService) GenerateRoundRobinSchedule() error {
	seasonTeams, err := a.repo.GetSeasonTeam()
	if err != nil {
		return err
	}

	seasonMap := groupTeamsBySeason(seasonTeams)

	for seasonID, teamIDs := range seasonMap {
		matches := generateMatchesForSeason(seasonID, teamIDs)

		if err := a.matchRepo.PostMatches(matches); err != nil {
			return err
		}
	}

	return nil
}

func groupTeamsBySeason(seasonTeams []domain.SeasonTeam) map[uuid.UUID][]uuid.UUID {
	seasonMap := make(map[uuid.UUID][]uuid.UUID)
	for _, st := range seasonTeams {
		seasonMap[st.SeasonID] = append(seasonMap[st.SeasonID], st.TeamID)
	}
	return seasonMap
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

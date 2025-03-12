package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/match"
	"github.com/robertobouses/online-football-tycoon/team"
)

func main() {
	homePlayers := []team.Player{
		{PlayerId: uuid.New(), FirstName: "Lionel", LastName: "Messi", Technique: 95, Mental: 90, Physique: 80},
		{PlayerId: uuid.New(), FirstName: "Cristiano", LastName: "Ronaldo", Technique: 92, Mental: 85, Physique: 88},
		{PlayerId: uuid.New(), FirstName: "Neymar", LastName: "Junior", Technique: 90, Mental: 80, Physique: 75},
	}

	awayPlayers := []team.Player{
		{PlayerId: uuid.New(), FirstName: "Kevin", LastName: "De Bruyne", Technique: 93, Mental: 88, Physique: 82},
		{PlayerId: uuid.New(), FirstName: "Luka", LastName: "Modric", Technique: 91, Mental: 89, Physique: 78},
		{PlayerId: uuid.New(), FirstName: "Erling", LastName: "Haaland", Technique: 85, Mental: 83, Physique: 95},
	}

	homeTeam := team.Team{Name: "FC Barcelona", Country: "Spain", Players: homePlayers}
	awayTeam := team.Team{Name: "Manchester City", Country: "England", Players: awayPlayers}

	homeStrategy := match.Strategy{StrategyTeam: homeTeam, Formation: "4-3-3", GameTempo: "fast"}
	awayStrategy := match.Strategy{StrategyTeam: awayTeam, Formation: "4-4-2", GameTempo: "normal"}

	game := match.Match{
		HomeMatchStrategy: homeStrategy,
		AwayMatchStrategy: awayStrategy,
	}

	result, err := game.Play()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Final Score:\n%s %d - %d %s\n", homeTeam.Name, result.HomeStats.Goals, result.AwayStats.Goals, awayTeam.Name)
}

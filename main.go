package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/robertobouses/online-football-tycoon/match"
	"github.com/robertobouses/online-football-tycoon/team"
)

func main() {

<<<<<<< HEAD
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Get("/play", func(w http.ResponseWriter, r *http.Request) {
		homePlayers := []team.Player{
			{PlayerId: uuid.New(), FirstName: "Marc-André", LastName: "ter Stegen", Nationality: "Germany", Position: "goalkeeper", Age: 31, Fee: 50000000, Salary: 10000000, Technique: 85, Mental: 88, Physique: 80, InjuryDays: 0, Lined: true, Familiarity: 90, Fitness: 95, Happiness: 90},
			{PlayerId: uuid.New(), FirstName: "Jules", LastName: "Koundé", Nationality: "France", Position: "defender", Age: 25, Fee: 60000000, Salary: 9000000, Technique: 78, Mental: 85, Physique: 88, InjuryDays: 0, Lined: true, Familiarity: 85, Fitness: 92, Happiness: 87},
			{PlayerId: uuid.New(), FirstName: "Ronald", LastName: "Araújo", Nationality: "Uruguay", Position: "defender", Age: 24, Fee: 70000000, Salary: 9500000, Technique: 80, Mental: 87, Physique: 90, InjuryDays: 0, Lined: true, Familiarity: 88, Fitness: 94, Happiness: 88},
			{PlayerId: uuid.New(), FirstName: "Andreas", LastName: "Christensen", Nationality: "Denmark", Position: "defender", Age: 27, Fee: 40000000, Salary: 8000000, Technique: 76, Mental: 85, Physique: 85, InjuryDays: 0, Lined: true, Familiarity: 86, Fitness: 91, Happiness: 85},
			{PlayerId: uuid.New(), FirstName: "Alejandro", LastName: "Balde", Nationality: "Spain", Position: "defender", Age: 20, Fee: 50000000, Salary: 7000000, Technique: 78, Mental: 80, Physique: 89, InjuryDays: 0, Lined: true, Familiarity: 83, Fitness: 95, Happiness: 89},
			{PlayerId: uuid.New(), FirstName: "Pedri", LastName: "González", Nationality: "Spain", Position: "midfielder", Age: 21, Fee: 100000000, Salary: 12000000, Technique: 92, Mental: 88, Physique: 78, InjuryDays: 0, Lined: true, Familiarity: 90, Fitness: 92, Happiness: 91},
			{PlayerId: uuid.New(), FirstName: "Frenkie", LastName: "de Jong", Nationality: "Netherlands", Position: "midfielder", Age: 26, Fee: 90000000, Salary: 11000000, Technique: 90, Mental: 87, Physique: 85, InjuryDays: 0, Lined: true, Familiarity: 89, Fitness: 91, Happiness: 90},
			{PlayerId: uuid.New(), FirstName: "Gavi", LastName: "Paez", Nationality: "Spain", Position: "midfielder", Age: 19, Fee: 80000000, Salary: 9000000, Technique: 88, Mental: 85, Physique: 80, InjuryDays: 0, Lined: true, Familiarity: 87, Fitness: 93, Happiness: 90},
			{PlayerId: uuid.New(), FirstName: "Raphinha", LastName: "Dias", Nationality: "Brazil", Position: "forward", Age: 27, Fee: 60000000, Salary: 10000000, Technique: 85, Mental: 82, Physique: 86, InjuryDays: 0, Lined: true, Familiarity: 85, Fitness: 92, Happiness: 88},
			{PlayerId: uuid.New(), FirstName: "Robert", LastName: "Lewandowski", Nationality: "Poland", Position: "forward", Age: 35, Fee: 50000000, Salary: 12000000, Technique: 92, Mental: 90, Physique: 88, InjuryDays: 0, Lined: true, Familiarity: 90, Fitness: 90, Happiness: 85},
			{PlayerId: uuid.New(), FirstName: "João", LastName: "Félix", Nationality: "Portugal", Position: "forward", Age: 24, Fee: 70000000, Salary: 9500000, Technique: 88, Mental: 83, Physique: 82, InjuryDays: 0, Lined: true, Familiarity: 87, Fitness: 92, Happiness: 86},
		}

		awayPlayers := []team.Player{
			{PlayerId: uuid.New(), FirstName: "Ederson", LastName: "Moraes", Nationality: "Brazil", Position: "goalkeeper", Age: 31, Fee: 60000000, Salary: 11000000, Technique: 86, Mental: 89, Physique: 85, InjuryDays: 0, Lined: true, Familiarity: 91, Fitness: 95, Happiness: 90},
			{PlayerId: uuid.New(), FirstName: "Kyle", LastName: "Walker", Nationality: "England", Position: "defender", Age: 34, Fee: 40000000, Salary: 9000000, Technique: 80, Mental: 87, Physique: 90, InjuryDays: 0, Lined: true, Familiarity: 85, Fitness: 93, Happiness: 87},
			{PlayerId: uuid.New(), FirstName: "Ruben", LastName: "Dias", Nationality: "Portugal", Position: "defender", Age: 26, Fee: 80000000, Salary: 10000000, Technique: 82, Mental: 89, Physique: 92, InjuryDays: 0, Lined: true, Familiarity: 88, Fitness: 94, Happiness: 88},
			{PlayerId: uuid.New(), FirstName: "John", LastName: "Stones", Nationality: "England", Position: "defender", Age: 29, Fee: 70000000, Salary: 9500000, Technique: 81, Mental: 88, Physique: 89, InjuryDays: 0, Lined: true, Familiarity: 87, Fitness: 92, Happiness: 87},
			{PlayerId: uuid.New(), FirstName: "Josko", LastName: "Gvardiol", Nationality: "Croatia", Position: "defender", Age: 22, Fee: 90000000, Salary: 10500000, Technique: 79, Mental: 85, Physique: 91, InjuryDays: 0, Lined: true, Familiarity: 86, Fitness: 93, Happiness: 89},
			{PlayerId: uuid.New(), FirstName: "Rodri", LastName: "Hernandez", Nationality: "Spain", Position: "midfielder", Age: 27, Fee: 100000000, Salary: 12000000, Technique: 91, Mental: 89, Physique: 85, InjuryDays: 0, Lined: true, Familiarity: 90, Fitness: 92, Happiness: 91},
			{PlayerId: uuid.New(), FirstName: "Kevin", LastName: "De Bruyne", Nationality: "Belgium", Position: "midfielder", Age: 33, Fee: 100000000, Salary: 13000000, Technique: 95, Mental: 92, Physique: 84, InjuryDays: 0, Lined: true, Familiarity: 92, Fitness: 90, Happiness: 90},
			{PlayerId: uuid.New(), FirstName: "Bernardo", LastName: "Silva", Nationality: "Portugal", Position: "midfielder", Age: 30, Fee: 80000000, Salary: 11000000, Technique: 89, Mental: 87, Physique: 80, InjuryDays: 0, Lined: true, Familiarity: 89, Fitness: 91, Happiness: 89},
			{PlayerId: uuid.New(), FirstName: "Phil", LastName: "Foden", Nationality: "England", Position: "forward", Age: 24, Fee: 90000000, Salary: 11500000, Technique: 88, Mental: 85, Physique: 82, InjuryDays: 0, Lined: true, Familiarity: 87, Fitness: 92, Happiness: 88},
			{PlayerId: uuid.New(), FirstName: "Erling", LastName: "Haaland", Nationality: "Norway", Position: "forward", Age: 24, Fee: 180000000, Salary: 15000000, Technique: 92, Mental: 90, Physique: 95, InjuryDays: 0, Lined: true, Familiarity: 90, Fitness: 90, Happiness: 90},
		}

		homeTeam := team.Team{Name: "FC Barcelona", Country: "Spain", Players: homePlayers}
		awayTeam := team.Team{Name: "Manchester City", Country: "England", Players: awayPlayers}

		homeStrategy := match.Strategy{
			StrategyTeam:         homeTeam,
			Formation:            "4-4-2",
			PlayingStyle:         "possession",
			GameTempo:            "balanced_tempo",
			PassingStyle:         "short",
			DefensivePositioning: "zonal_marking",
			BuildUpPlay:          "play_from_back",
			AttackFocus:          "wide_play",
			KeyPlayerUsage:       "reference_player",
		}
		awayStrategy := match.Strategy{
			StrategyTeam:         awayTeam,
			Formation:            "4-4-2",
			PlayingStyle:         "possession",
			GameTempo:            "balanced_tempo",
			PassingStyle:         "short",
			DefensivePositioning: "zonal_marking",
			BuildUpPlay:          "play_from_back",
			AttackFocus:          "wide_play",
			KeyPlayerUsage:       "reference_player"}

		game := match.Match{
			HomeMatchStrategy: homeStrategy,
			AwayMatchStrategy: awayStrategy,
		}

		result, err := game.Play()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res := response{
			Result: result,
			Match:  game,
		}
		responseBody, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(responseBody)
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
=======
	homePlayers := []team.Player{
		{PlayerId: uuid.New(), FirstName: "Marc-André", LastName: "ter Stegen", Nationality: "Germany", Position: "goalkeeper", Age: 31, Fee: 50000000, Salary: 10000000, Technique: 85, Mental: 88, Physique: 80, InjuryDays: 0, Lined: true, Familiarity: 90, Fitness: 95, Happiness: 90},
		{PlayerId: uuid.New(), FirstName: "Jules", LastName: "Koundé", Nationality: "France", Position: "defender", Age: 25, Fee: 60000000, Salary: 9000000, Technique: 78, Mental: 85, Physique: 88, InjuryDays: 0, Lined: true, Familiarity: 85, Fitness: 92, Happiness: 87},
		{PlayerId: uuid.New(), FirstName: "Ronald", LastName: "Araújo", Nationality: "Uruguay", Position: "defender", Age: 24, Fee: 70000000, Salary: 9500000, Technique: 80, Mental: 87, Physique: 90, InjuryDays: 0, Lined: true, Familiarity: 88, Fitness: 94, Happiness: 88},
		{PlayerId: uuid.New(), FirstName: "Andreas", LastName: "Christensen", Nationality: "Denmark", Position: "defender", Age: 27, Fee: 40000000, Salary: 8000000, Technique: 76, Mental: 85, Physique: 85, InjuryDays: 0, Lined: true, Familiarity: 86, Fitness: 91, Happiness: 85},
		{PlayerId: uuid.New(), FirstName: "Alejandro", LastName: "Balde", Nationality: "Spain", Position: "defender", Age: 20, Fee: 50000000, Salary: 7000000, Technique: 78, Mental: 80, Physique: 89, InjuryDays: 0, Lined: true, Familiarity: 83, Fitness: 95, Happiness: 89},
		{PlayerId: uuid.New(), FirstName: "Pedri", LastName: "González", Nationality: "Spain", Position: "midfielder", Age: 21, Fee: 100000000, Salary: 12000000, Technique: 92, Mental: 88, Physique: 78, InjuryDays: 0, Lined: true, Familiarity: 90, Fitness: 92, Happiness: 91},
		{PlayerId: uuid.New(), FirstName: "Frenkie", LastName: "de Jong", Nationality: "Netherlands", Position: "midfielder", Age: 26, Fee: 90000000, Salary: 11000000, Technique: 90, Mental: 87, Physique: 85, InjuryDays: 0, Lined: true, Familiarity: 89, Fitness: 91, Happiness: 90},
		{PlayerId: uuid.New(), FirstName: "Gavi", LastName: "Paez", Nationality: "Spain", Position: "midfielder", Age: 19, Fee: 80000000, Salary: 9000000, Technique: 88, Mental: 85, Physique: 80, InjuryDays: 0, Lined: true, Familiarity: 87, Fitness: 93, Happiness: 90},
		{PlayerId: uuid.New(), FirstName: "Raphinha", LastName: "Dias", Nationality: "Brazil", Position: "forward", Age: 27, Fee: 60000000, Salary: 10000000, Technique: 85, Mental: 82, Physique: 86, InjuryDays: 0, Lined: true, Familiarity: 85, Fitness: 92, Happiness: 88},
		{PlayerId: uuid.New(), FirstName: "Robert", LastName: "Lewandowski", Nationality: "Poland", Position: "forward", Age: 35, Fee: 50000000, Salary: 12000000, Technique: 92, Mental: 90, Physique: 88, InjuryDays: 0, Lined: true, Familiarity: 90, Fitness: 90, Happiness: 85},
		{PlayerId: uuid.New(), FirstName: "João", LastName: "Félix", Nationality: "Portugal", Position: "forward", Age: 24, Fee: 70000000, Salary: 9500000, Technique: 88, Mental: 83, Physique: 82, InjuryDays: 0, Lined: true, Familiarity: 87, Fitness: 92, Happiness: 86},
>>>>>>> 4e7b465 (add get matches test)
	}
	log.Println("Starting server on port ", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)

}

type response struct {
	Match  match.Match `json:"match"`
	Result match.Result
}

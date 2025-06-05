package player

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/robertobouses/online-football-tycoon/internal/domain"
)

func (a AppService) GeneratePlayer(country, position string) (domain.Player, error) {

	firstName, lastName, err := GetRandomNameByCountry(country)
	if err != nil {
		return domain.Player{}, fmt.Errorf("error generating player name: %v", err)
	}

	age, technique, mental, physique, injuryDays := CalculatePlayerAtributes()
	fee, salary := CalculatePlayerFeeAndSalary(technique, mental, physique, age, country, position)

	player := domain.Player{
		FirstName:   firstName,
		LastName:    lastName,
		Nationality: country,
		Position:    position,
		Age:         age,
		Fee:         fee,
		Salary:      salary,
		Technique:   technique,
		Mental:      mental,
		Physique:    physique,
		InjuryDays:  injuryDays,
		Lined:       false,
		Familiarity: rand.Intn(80) + 1,
		Fitness:     rand.Intn(100) + 1,
		Happiness:   rand.Intn(50) + 1,
	}
	log.Printf("Generated player: %+v\n", player)

	err = a.repo.PostPlayer(player)
	if err != nil {
		log.Printf("Error saving player to repository: %v\n", err)
		return domain.Player{}, fmt.Errorf("error saving player: %w", err)
	}
	log.Println("Player successfully saved to repository")

	return player, nil
}

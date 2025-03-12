package main

import (
	"fmt"

	"github.com/robertobouses/online-football-tycoon/match"
)

func main() {
	fmt.Println("Online Football Tycoon")

	m := match.Match{
		HomeMatchStrategy: match.Strategy{
			GameTempo: "fast",
		},
		AwayMatchStrategy: match.Strategy{
			GameTempo: "slow",
		},
	}

	result, _ := m.Play()
	fmt.Println(result)

}

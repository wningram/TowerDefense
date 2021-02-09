package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello WOrld\n")
	tick := time.Tick(1 * time.Second)
	players := map[Location]*Player{
		Location{80, 2}: &Player{true, true, true},
		Location{90, 5}: &Player{true, true, true},
		Location{95, 5}: &Player{true, true, false},
	}

	bf := BattleField{
		Length:         100,
		Height:         10,
		DefenseLineLoc: 90,
		Players:        players,
	}
	kb := Keyboard{&players}

	// Listen for keyboard commands
	go kb.BeginListening()

	for {
		select {
		case <-tick:
			bf.Next()
		default:
			continue
		}
	}
}

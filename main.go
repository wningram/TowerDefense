package main

import (
	"time"
)

func main() {
	tick := time.Tick(500 * time.Millisecond)
	players := map[Location]*Player{
		{80, 2}: {true, true, true},
		{90, 5}: {true, true, true},
		{95, 5}: {true, true, false},
		{85, 5}: {true, true, true},
	}

	bf := BattleField{
		Length:         100,
		Height:         10,
		DefenseLineLoc: 90,
		Players:        players,
		MaxEnemies:     5,
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

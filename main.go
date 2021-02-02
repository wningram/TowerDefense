package main

import "fmt"

func main() {
	fmt.Printf("Hello WOrld\n")
	players := map[Location]*Player{
		Location{10, 10}: &Player{true, true},
		Location{90, 5}:  &Player{true, true},
	}

	bf := BattleField{
		Length:         100,
		Height:         50,
		DefenseLineLoc: 90,
		Players:        players,
	}

	bf.Draw()
}

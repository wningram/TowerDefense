package main

import "fmt"

type Stats struct {
	Bf BattleField
}

func (s Stats) Print() {
	players := s.Bf.Players
	for loc, player := range players {
		if !player.Bot {
			fmt.Printf("UserPlayer Y: %d\n", loc.Y)
			fmt.Printf("UserPlayer X: %d\n", loc.X)
			fmt.Printf("Total enemies: %d\n", s.Bf.CountEnemies())
			fmt.Print("---------------------------\n")
			return
		}
	}
}

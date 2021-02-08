package main

import (
	"ansi"
	"fmt"
)

type BattleField struct {
	Length         int
	Height         int
	DefenseLineLoc int
	Players        map[Location]*Player
}

// Draw draws the battlefield to the console, including all of the players.
func (bf BattleField) Draw() error {
	if bf.Length < 1 {
		return fmt.Errorf("Cannot draw battlefield with no length.")
	}

	if bf.Height < 0 {
		return fmt.Errorf("Cannot draw battlefield with no height.")
	}

	if bf.DefenseLineLoc < 1 {
		return fmt.Errorf("Cannot draw defense line at beginning of battlefield.")
	}

	for y := 0; y < bf.Height; y++ {
		for x := 0; x < bf.Length; x++ {
			// If we are at the defense line and no player exists here, draw the wall
			if x == bf.DefenseLineLoc {
				if _, exists := bf.Players[Location{x, y}]; !exists {
					fmt.Printf("|")
				} else if player, exists := bf.Players[Location{x, y}]; exists && player.Enemy {
					fmt.Printf(" ")
					// Deactive player so that it does not move forward anymore
					bf.Players[Location{x, y}].Active = false
				}
			} else if player, exists := bf.Players[Location{x, y}]; exists && player.Active {
				if !player.Bot {
					fmt.Printf("=|")
				} else if player.Enemy {
					fmt.Printf("*")
				} else {
					fmt.Printf("#")
				}
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println("-----------------------\n")
	return nil
}

// Erase erases the console output representation of the battlefield.
func (bf BattleField) Erase() {
	for y := bf.Height; y > 0; y-- {
		fmt.Print(ansi.DL)
		fmt.Print(ansi.CUU)
	}
}

// Next moves each player in its specified direction by one increment.
func (bf BattleField) Next() {
	for loc, player := range bf.Players {
		// DOn't increment X value for UserPlayer
		if !player.Bot {
			continue
		}

		// If the next space over from an enemy is an ally, then delete the current enemy
		if player, exists := bf.Players[Location{loc.X + 1, loc.Y}]; exists && !player.Enemy {
			delete(bf.Players, loc)
			continue
		}

		if loc.X != bf.DefenseLineLoc && player.Enemy {
			delete(bf.Players, loc)
			loc.X++
			bf.Players[loc] = player
		} else if !player.Enemy {
			delete(bf.Players, loc)
			loc.X--
			bf.Players[loc] = player
		}
	}
	bf.Draw()
	stats := Stats{bf}
	stats.Print()
}

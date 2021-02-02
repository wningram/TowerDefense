package main

import "fmt"

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
				} else {
					fmt.Printf(" ")
					// Deactive player so that it does not move forward anymore
					bf.Players[Location{x, y}].Active = false
				}
			} else if player, exists := bf.Players[Location{x, y}]; exists && player.Active {
				if player.Enemy {
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
	fmt.Println()
	return nil
}

// Erase erases the console output representation of the battlefield.
func (bf BattleField) Erase() error {
	return fmt.Errorf("Not yet implemented.")
}

// TODO: Add function to move player forward, should be triggered every iteration

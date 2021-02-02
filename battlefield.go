package main

import "fmt"

type BattleField struct {
	Length         int
	DefenseLineLoc int
	Players        []Player
}

// Draw draws the battlefield to the console, including all of the players.
func (bf BattleField) Draw() {
	fmt.Printf("Function not yet implemented.")
}

// Erase erases the console output representation of the battlefield.
func (bf BattleField) Erase() {
	fmt.Printf("Function not yet implemented.")
}

package main

import (
	"fmt"
	"testing"
)

func TestMoveUserPlayer(t *testing.T) {
	// Setup
	userPlayer := Player{false, true, false}
	players := map[Location]*Player{
		{95, 5}: &userPlayer,
	}
	kb := Keyboard{&players}

	t.Run("MoveUp", func(t *testing.T) {
		// Execution
		kb.moveUserPlayer(1)

		// Assertions
		passUp := false
		for loc, player := range players {
			// Location should be (95, 6)
			if loc.Y == 6 && !player.Bot {
				passUp = true
				break
			}
		}
		if !passUp {
			t.Fail()
		}
	})

	t.Run("MoveDown", func(t *testing.T) {
		// Execution
		kb.moveUserPlayer(-1)

		// Assertions
		passDown := false
		for loc, player := range players {
			// Location should be (95, 5)
			if loc.Y == 5 && !player.Bot {
				passDown = true
				break
			}
		}
		if !passDown {
			t.Fail()
		}
	})

	t.Run("UserPlayerNotFound", func(t *testing.T) {
		userPlayer.Bot = true

		// Execution
		result := kb.moveUserPlayer(1)
		if _, ok := result.(error); !ok {
			t.Fail()
		}
	})
}

func TestAddAllyPlayer(t *testing.T) {
	// Setup
	userPlayer := Player{false, true, false}
	players := map[Location]*Player{
		{95, 5}: &userPlayer,
	}
	kb := Keyboard{&players}

	// Execution
	kb.addAllyPlayer()

	// Assertions
	fmt.Printf("players length: %d\n", len(players))
	if len(players) != 2 {
		t.Fail()
	}
}

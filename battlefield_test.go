package main

import "testing"

func TestCoutnEnemies(t *testing.T) {
	player1 := Player{Enemy: true, Active: true, Bot: true}  // Should count
	player2 := Player{Enemy: true, Active: true, Bot: true}  // Should count
	player3 := Player{Enemy: false, Active: true, Bot: true} // Should not count
	player4 := Player{Enemy: true, Active: false, Bot: true} // Should not count
	bf := BattleField{
		Length:         10,
		Height:         10,
		DefenseLineLoc: 7,
		Players: map[Location]*Player{
			{1, 1}: &player1,
			{1, 2}: &player2,
			{1, 3}: &player3,
			{1, 4}: &player4,
		},
		MaxEnemies: 10,
	}

	// Execution
	enemyCount := bf.CountEnemies()

	// Assertions
	if enemyCount != 2 {
		t.Fail()
	}
}

func TestInjectEnemies(t *testing.T) {
	player1 := Player{Enemy: true, Active: true, Bot: true}
	player2 := Player{Enemy: true, Active: true, Bot: true}
	player3 := Player{Enemy: true, Active: true, Bot: true}
	bf := BattleField{
		Length:         10,
		Height:         10,
		DefenseLineLoc: 7,
		Players: map[Location]*Player{
			{1, 1}: &player1,
			{1, 2}: &player2,
			{1, 3}: &player3,
		},
		MaxEnemies: 2,
	}
	t.Run("OverMaxLimit", func(t *testing.T) {
		bf.InjectEnemies()
		if bf.CountEnemies() != 3 {
			t.Fail()
		}
	})

	t.Run("UnderMaxLimit", func(t *testing.T) {
		bf.MaxEnemies = 5
		bf.InjectEnemies()
		if bf.CountEnemies() != 5 {
			t.Fail()
		}
	})
}

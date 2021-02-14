package main

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

type Keyboard struct {
	PlayersPtr *map[Location]*Player
}

func (kb *Keyboard) moveUserPlayer(velocity int) error {
	for loc, player := range *kb.PlayersPtr {
		if !player.Bot {
			delete(*kb.PlayersPtr, loc)
			loc.Y += velocity
			(*kb.PlayersPtr)[loc] = player
			return nil
		}
	}
	return fmt.Errorf("Could not find UserPlayer.")
}

func (kb *Keyboard) addAllyPlayer() {
	var (
		xPos int
		yPos int
	)
	for loc, player := range *kb.PlayersPtr {
		if !player.Bot { // If we have iterated to UserPlayer
			// Set the position of the new ally to one space left of the UserPlayer
			xPos = loc.X - 1
			yPos = loc.Y
			break
		}
	}
	// Add ally player to players map
	(*kb.PlayersPtr)[Location{xPos, yPos}] = &Player{false, true, true}
}

func (kb *Keyboard) BeginListening() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	eventQueue := make(chan termbox.Event)
	// Begin listening for key events
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyArrowUp:
					kb.moveUserPlayer(-1)
				case termbox.KeyArrowDown:
					kb.moveUserPlayer(1)
				case termbox.KeySpace:
					kb.addAllyPlayer()
				case termbox.KeyCtrlC:
					// Stop listening for keyboard events
					termbox.Close()
					os.Exit(0)
				}
			} else if ev.Type == termbox.EventInterrupt {
				// Stop listening for keyboard events
				termbox.Close()
				os.Exit(0)
			}
		default:
			continue
			//time.Sleep(tickDuration * time.Millisecond)
		}
	}
}

// Game Struct and Functions

package main

import "time"

type Game struct {
	Board       *Board
	Screen      string
	State       string
	StartTime   time.Time
	ElapsedTime time.Duration
	EndTime     time.Time
}

// Initialize a Game

func NewGame() *Game {

	return &Game{
		Screen: "Menu",
		State:  "Menu",
	}
}

// Start a Game

func (g *Game) StartGame(Size int, MineCount int) {

	g.Board = NewBoard(Size, MineCount)

	g.Screen = "Minefield"
	g.State = "Playing"
	g.StartTime = time.Now()

}

// Update the timer

func (g *Game) UpdateTimer() {
	if g.State == "Playing" {
		g.ElapsedTime = time.Since(g.StartTime)
	}
}

// Check the status of the Game

func (g *Game) CheckStatus() {

	if g.Board.isWon {

		g.State = "Won"

	} else if g.Board.isLost {

		g.State = "Lost"
	}

}

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

func NewGame() *Game {

	return &Game{
		Screen: "Menu",
		State:  "Menu",
	}
}

func (g *Game) StartGame() {

	var Size int = 16
	var MineCount int = 5

	g.Board = NewBoard(Size, MineCount)

	g.Screen = "Minefield"
	g.State = "Playing"
	g.StartTime = time.Now()

}

func (g *Game) UpdateTimer() {

	g.ElapsedTime = time.Since(g.StartTime)

}

func (g *Game) CheckStatus() {

	if g.Board.isWon {

		g.State = "Won"

	} else if g.Board.isLost {

		g.State = "Lost"
	}

}

// Initialization of the game field

package main

// Cell object

type Cell struct {
	isMine        bool
	isRevealed    bool
	isFlagged     bool
	adjacentBombs int
}

func (c *Cell) SetRevealed() {

	c.isRevealed = true

}

func (c *Cell) ToggleFlagged() {

	if !c.isRevealed {
		c.isFlagged = !c.isFlagged
	}

}

// Cell Struct and Functions

package main

// Cell object

type Cell struct {
	isMine        bool
	isRevealed    bool
	isFlagged     bool
	adjacentBombs int
}

// Set cell to revealed

func (c *Cell) SetRevealed() {

	c.isRevealed = true

}

// Set cell to flagged

func (c *Cell) ToggleFlagged() {

	c.isFlagged = !c.isFlagged

}

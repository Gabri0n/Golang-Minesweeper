// Game Functions

package main

import (
	"fmt"
	"math/rand/v2"
)

type Board struct {
	Cells     [][]Cell
	Size      int
	MineCount int
}

// Build a board

func NewBoard(Size, MineCount int) *Board {

	Cells := make([][]Cell, Size)

	for i := range Cells {
		Cells[i] = make([]Cell, Size)
	}

	Board := &Board{
		Cells:     Cells,
		Size:      Size,
		MineCount: MineCount,
	}

	Board.AddMines(MineCount, Board.Cells)
	Board.CountMines(Board.Cells)

	return Board
}

func (b *Board) Select(x int, y int) {

	switch {
	case b.Cells[x][y].isFlagged:

		// Do Nothing

	case b.Cells[x][y].isMine:

		b.MineSelected(b.Cells)

	case !b.Cells[x][y].isFlagged && !b.Cells[x][y].isRevealed:

		b.FloodFill(x, y)

	}

}

// Adds bombs randomly to the field

func (b *Board) AddMines(bomb_count int, cell [][]Cell) [][]Cell {

	fmt.Print("Bomb count = ", bomb_count)
	fmt.Println()

	for count := 0; count < bomb_count; {
		var x = rand.IntN(len(cell))
		var y = rand.IntN(len(cell))

		if cell[x][y].isMine == false {
			cell[x][y].isMine = true
			count++
		}
	}

	return cell
}

// Find the adjacent bomb count for each cell

func (b *Board) CountMines(cell [][]Cell) [][]Cell {

	var direction_x = []int{-1, -1, -1, 0, 0, 1, 1, 1}
	var direction_y = []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for x := 0; x < len(cell); x++ {
		for y := 0; y < len(cell[0]); y++ {
			if cell[x][y].isMine == false {

				var count int = 0

				for i := 0; i < 8; i++ {
					new_x := x + direction_x[i]
					new_y := y + direction_y[i]

					if new_x >= 0 && new_x < len(cell) && new_y >= 0 && new_y < len(cell[0]) {
						if cell[new_x][new_y].isMine {
							count++
						}
					}
				}

				cell[x][y].adjacentBombs = count
				count = 0
			}
		}
	}

	return cell

}

// Reveal all bombs on selection of one

func (b *Board) MineSelected(cell [][]Cell) [][]Cell {

	for x := 0; x < len(cell); x++ {
		for y := 0; y < len(cell[0]); y++ {
			if cell[x][y].isMine {
				cell[x][y].isRevealed = true
			}

		}

	}

	return cell
}

func (b *Board) FloodFill(xpos int, ypos int) [][]Cell {

	var direction_x = []int{-1, -1, -1, 0, 0, 1, 1, 1}
	var direction_y = []int{-1, 0, 1, -1, 1, -1, 0, 1}

	type point struct {
		queue_x, queue_y int
	}

	b.Cells[xpos][ypos].isRevealed = true
	queue := []point{{xpos, ypos}}

	for front := 0; front < len(queue); front++ {
		target := queue[front]

		for i := 0; i < 8; i++ {
			new_x := target.queue_x + direction_x[i]
			new_y := target.queue_y + direction_y[i]

			if new_x < 0 || new_x >= len(b.Cells) || new_y < 0 || new_y >= len(b.Cells[0]) {
				continue
			}

			if b.Cells[new_x][new_y].isMine || b.Cells[new_x][new_y].isRevealed {
				continue
			}

			if b.Cells[new_x][new_y].adjacentBombs > 0 {
				b.Cells[new_x][new_y].isRevealed = true
				continue
			}

			b.Cells[new_x][new_y].isRevealed = true

			queue = append(queue, point{new_x, new_y})

		}

	}
	return b.Cells
}

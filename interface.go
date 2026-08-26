// User interface

package golang_minesweeper

import (
	"fmt"
)

// Render field for debugging

func render_field(matrix [][]Cell) {
	for _, row := range matrix {
		for _, val := range row {
			switch {
			case val.bomb == true:
				fmt.Print("🟥") // Bomb

			case val.revealed == true && val.adjacentBombs > 0:
				fmt.Printf("%d ", val.adjacentBombs)

			case val.revealed == true:
				fmt.Print("⬛") // Revealed

			default:
				fmt.Print("⬜")
			}
		}
		fmt.Println()
	}
}

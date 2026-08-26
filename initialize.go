// Initialization of the game field

package main

import (
	"fmt"
	"math/rand/v2"
)

type Cell struct {
	bomb          bool
	revealed      bool
	flagged       bool
	adjacentBombs int
}

func initialize_field(bomb_count int, matrix_size int) [][]Cell {

	matrix := count_bombs(add_bombs(bomb_count, build_field(matrix_size)))

	return matrix
}

// Builds the field from the supplied size

func build_field(matrix_size int) [][]Cell {

	matrix := make([][]Cell, matrix_size)

	for i := range matrix {
		matrix[i] = make([]Cell, matrix_size)
	}

	return matrix
}

// Adds bombs randomly to the field

func add_bombs(bomb_count int, matrix [][]Cell) [][]Cell {

	fmt.Print("Bomb count = ", bomb_count)
	fmt.Println()

	for count := 0; count < bomb_count; {
		var x = rand.IntN(len(matrix))
		var y = rand.IntN(len(matrix))

		if matrix[x][y].bomb == false {
			matrix[x][y].bomb = true
			count++
		}
	}

	return matrix
}

// Find the adjacent bomb count for each cell

func count_bombs(matrix [][]Cell) [][]Cell {

	var direction_x = []int{-1, -1, -1, 0, 0, 1, 1, 1}
	var direction_y = []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			if matrix[x][y].bomb == false {

				var count int = 0

				for i := 0; i < 8; i++ {
					new_x := x + direction_x[i]
					new_y := y + direction_y[i]

					if new_x >= 0 && new_x < len(matrix) && new_y >= 0 && new_y < len(matrix[0]) {
						if matrix[new_x][new_y].bomb {
							count++
						}
					}
				}

				matrix[x][y].adjacentBombs = count
				count = 0
			}
		}
	}

	return matrix

}

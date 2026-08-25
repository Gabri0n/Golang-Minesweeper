package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	var matrix_size int = 25
	var bomb_count int = 30

	matrix := add_bombs(bomb_count, build_field(matrix_size))
	matrix = count_bombs(matrix)
	render_field(matrix)

}

func build_field(matrix_size int) [][]int {

	matrix := make([][]int, matrix_size)

	for i := range matrix {
		matrix[i] = make([]int, matrix_size)
	}

	return matrix
}

func add_bombs(bomb_count int, matrix [][]int) [][]int {

	fmt.Print("Bomb count = ", bomb_count)
	fmt.Println()

	for count := 0; count < bomb_count; {
		var x = rand.IntN(len(matrix))
		var y = rand.IntN(len(matrix))

		if matrix[x][y] == 0 {
			matrix[x][y] = -4
			count++
		}
	}

	return matrix
}

func count_bombs(matrix [][]int) [][]int {

	var direction_x = []int{-1, -1, -1, 0, 0, 1, 1, 1}
	var direction_y = []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			if matrix[x][y] == 0 {

				var count int = 0

				for i := 0; i < 8; i++ {
					new_x := x + direction_x[i]
					new_y := y + direction_y[i]

					if new_x >= 0 && new_x < len(matrix) && new_y >= 0 && new_y < len(matrix[0]) {
						if matrix[new_x][new_y] == -4 {
							count++
						}

					}

				}

				matrix[x][y] = count
				count = 0
			}
		}

	}

	return matrix

}

func render_field(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			switch {
			case val == -4:
				fmt.Print("🟥") // Bomb
			case val == -3:
				fmt.Print("🚩") // Flagged and Bomb
			case val == -2:
				fmt.Print("🚩") // Flagged no Bomb
			case val == -1:
				fmt.Print("⬜") // Revealed
			case val == 0:
				fmt.Print("🔲") // Not Revealed
			default:
				fmt.Printf("%d ", val)
			}
		}
		fmt.Println()
	}
}

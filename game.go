// Game Functions

package main

// Flood fill alogrithm when selecting empty cell

func flood_fill(matrix [][]Cell, xpos int, ypos int) [][]Cell {

	var direction_x = []int{-1, -1, -1, 0, 0, 1, 1, 1}
	var direction_y = []int{-1, 0, 1, -1, 1, -1, 0, 1}

	type point struct {
		queue_x, queue_y int
	}

	matrix[xpos][ypos].revealed = true
	queue := []point{{xpos, ypos}}

	for front := 0; front < len(queue); front++ {
		target := queue[front]

		for i := 0; i < 8; i++ {
			new_x := target.queue_x + direction_x[i]
			new_y := target.queue_y + direction_y[i]

			if new_x < 0 || new_x >= len(matrix) || new_y < 0 || new_y >= len(matrix[0]) {
				continue
			}

			if matrix[new_x][new_y].bomb || matrix[new_x][new_y].revealed {
				continue
			}

			if matrix[new_x][new_y].adjacentBombs > 0 {
				matrix[new_x][new_y].revealed = true
				continue
			}

			matrix[new_x][new_y].revealed = true

			queue = append(queue, point{new_x, new_y})

		}

	}
	return matrix
}

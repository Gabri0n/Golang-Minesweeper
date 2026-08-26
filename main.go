// Main Program

package main

// Main Function, calls all other functions

func main() {

	var matrix_size int = 25
	var bomb_count int = 50

	matrix := initialize_field(bomb_count, matrix_size)
	// render_field(matrix)
	// matrix = flood_fill(matrix, 15, 15)
	// render_field(matrix)

	render_tui(matrix)
}

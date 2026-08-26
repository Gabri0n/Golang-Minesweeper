// Main Program

package main

// Main Function, calls all other functions

func main() {

	var matrix_size int = 50
	var bomb_count int = 250

	matrix := initialize_field(bomb_count, matrix_size)

	render_tui(matrix)
}

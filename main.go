// Main Program

package main

// Main Function, calls all other functions

func main() {

	var Size int = 50
	var MineCount int = 250

	b := NewBoard(Size, MineCount)

	RenderTui(b)
}

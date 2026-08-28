// Main Program

package main

// Main Function, calls all other functions

func main() {

	//var Size int = 50
	//var MineCount int = 250

	g := NewGame()
	g.StartGame()

	RenderTui(g)
}

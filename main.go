package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	size := 16
	mineCount := 30

	args := os.Args[1:]

	if len(args) >= 1 {
		val, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: size must be an integer")
			os.Exit(1)
		}
		size = val
	}

	if len(args) >= 2 {
		val, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: mine count must be an integer")
			os.Exit(1)
		}
		mineCount = val
	}

	g := NewGame()
	g.StartGame(size, mineCount)

	RenderTui(g)
}

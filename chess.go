package main

import (
	"fmt"
	"github.com/gnarlyskier/chess/game"
)

func main() {
	board := game.CreateStandardBoard()
	fmt.Println(board)
}

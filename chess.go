package main

import (
	"fmt"
	"github.com/gnarlyskier/chess/core"
)

func main() {
	board := core.CreateStandardBoard()
	fmt.Println(board)
}

package main

import (
	"fmt"
	"github.com/gnarlyskier/chess/token"
)

func main() {
	board := token.CreateStandardBoard()
	fmt.Println(board)
}

package main

import (
	"fmt"
	"github.com/gnarlyskier/chess/core"
	"github.com/gnarlyskier/chess/parse"
	"github.com/gnarlyskier/chess/token"
)

func main() {
	board := token.CreateStandardBoard()
	var player, otherPlayer core.Player = core.WHITE, core.BLACK
	for {
		fmt.Println(board)
		fmt.Printf("%s's move\n", player)
		move := parse.PromptMove(board, player)

		// Pretty print move
		status := fmt.Sprintf("%s: %s to %s", player, board.GetPos(move.Orig), move.Final.AsAlgPos())
		if capture := board.GetPos(move.Final); capture != nil {
			status += fmt.Sprintf(" (captured %s)", capture)
		}
		fmt.Println(status)
		fmt.Println("")

		board.DoMove(move)
		player, otherPlayer = otherPlayer, player
	}
}

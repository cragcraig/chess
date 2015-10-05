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
        move := parse.ReadMove(board, player)
        fmt.Printf("%s: %s to %s\n\n", player, board.GetPos(move.Orig), move.Final.AsAlgPos())
        board.DoMove(move)
        player, otherPlayer = otherPlayer, player
    }
}

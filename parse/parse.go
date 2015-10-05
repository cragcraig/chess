package parse

import (
	"bufio"
	"fmt"
	"github.com/gnarlyskier/chess/core"
	"github.com/gnarlyskier/chess/token"
	"os"
	"strings"
)

func readPos(prompt string) core.Position {
	bio := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, err := bio.ReadString('\n')
		if err != nil {
			continue
		}
		line = strings.TrimSpace(line)
		orig, err := core.AlgPos(line)
		if err == nil {
			return orig
		}
		fmt.Printf("Bad position: %s\n", line)
	}
}

func ReadMove(b token.Board, player core.Player) core.Move {
	for {
		move := core.Move{readPos("from> "), readPos("to> ")}
		if b.IsValidMove(move, player) {
			return move
		}
		fmt.Printf("Bad move: %s to %s\n", move.Orig.AsAlgPos(), move.Final.AsAlgPos())
	}
}

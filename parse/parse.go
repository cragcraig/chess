package parse

import (
	"bufio"
	"fmt"
	"github.com/gnarlyskier/chess/core"
	"github.com/gnarlyskier/chess/token"
	"os"
	"strings"
)

func promptPos(prompt string) core.Position {
	bio := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		if line, err := bio.ReadString('\n'); err == nil {
			line = strings.TrimSpace(line)
			if orig, err := core.AlgPos(line); err == nil {
				return orig
			}
			fmt.Printf("Failed to parse position: %s\n", line)
		}
	}
}

func PromptMove(b token.Board, player core.Player) core.Move {
	for {
		move := core.Move{promptPos("from> "), promptPos("to> ")}
		if b.IsValidMove(move, player) {
			return move
		}
		fmt.Printf("Invalid move: %s to %s\n", move.Orig.AsAlgPos(), move.Final.AsAlgPos())
	}
}

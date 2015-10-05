package game

import (
    "github.com/gnarlyskier/chess/core"
    "github.com/gnarlyskier/chess/token"
)

type Game struct {
	Moves []core.Move
	Start token.Board
}

func (g *Game) Add(m core.Move) {
	g.Moves = append(g.Moves, m)
}

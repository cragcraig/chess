package game

import (
    "github.com/gnarlyskier/chess/core"
)

type Game struct {
	Moves []core.Move
	Start Board
}

func (g *Game) Add(m core.Move) {
	g.Moves = append(g.Moves, m)
}

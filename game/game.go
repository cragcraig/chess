package game

import (
	"github.com/cragcraig/chess/core"
	"github.com/cragcraig/chess/token"
)

type Game struct {
	Moves []core.Move
	Start token.Board
}

func (g *Game) Add(m core.Move) {
	g.Moves = append(g.Moves, m)
}

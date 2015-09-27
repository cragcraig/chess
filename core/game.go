package core

type Game struct {
	Moves []Move
	Start Board
}

func (g *Game) Add(m Move) {
	g.Moves = append(g.Moves, m)
}

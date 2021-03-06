package token

import (
	"github.com/cragcraig/chess/core"
)

type bishop struct {
	protoToken
}

func (t *bishop) GetMoves(curPos core.Position, b *Board) []core.Position {
	pos := []core.Position{}
	pos = append(pos, b.extendMove(curPos, core.Offset{1, 1}, t.GetPlayer())...)
	pos = append(pos, b.extendMove(curPos, core.Offset{1, -1}, t.GetPlayer())...)
	pos = append(pos, b.extendMove(curPos, core.Offset{-1, 1}, t.GetPlayer())...)
	pos = append(pos, b.extendMove(curPos, core.Offset{-1, -1}, t.GetPlayer())...)
	return pos
}

func Bishop(p core.Player) Token {
	return &bishop{protoToken{p, 0, 'B', "bishop", -1}}
}

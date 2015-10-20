package token

import (
	"github.com/gnarlyskier/chess/core"
)

type queen struct {
	protoToken
}

func (t *queen) GetMoves(curPos core.Position, b Board) []core.Position {
	pos := []core.Position{}
	pos = append(pos, b.extendMove(curPos, core.Offset{1, 1}, t.GetPlayer())...)
	pos = append(pos, b.extendMove(curPos, core.Offset{1, -1}, t.GetPlayer())...)
	pos = append(pos, b.extendMove(curPos, core.Offset{-1, 1}, t.GetPlayer())...)
	pos = append(pos, b.extendMove(curPos, core.Offset{-1, -1}, t.GetPlayer())...)
	pos = append(pos, b.extendMove(curPos, core.Offset{1, 0}, t.GetPlayer())...)
	pos = append(pos, b.extendMove(curPos, core.Offset{0, 1}, t.GetPlayer())...)
	pos = append(pos, b.extendMove(curPos, core.Offset{-1, 0}, t.GetPlayer())...)
	pos = append(pos, b.extendMove(curPos, core.Offset{0, -1}, t.GetPlayer())...)
	return pos
}

func Queen(p core.Player) Token {
	return &queen{protoToken{p, 0, 'Q', "queen"}}
}

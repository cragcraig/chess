package token

import (
	"github.com/gnarlyskier/chess/core"
)

type rook struct {
	protoToken
}

func (t *rook) GetMoves(curPos core.Position, b Board) []core.Position {
	pos := []core.Position{}
	pos = append(pos, b.extendMove(curPos, core.Offset{1, 0}, t.GetPlayer())...)
	pos = append(pos, b.extendMove(curPos, core.Offset{0, 1}, t.GetPlayer())...)
	pos = append(pos, b.extendMove(curPos, core.Offset{-1, 0}, t.GetPlayer())...)
	pos = append(pos, b.extendMove(curPos, core.Offset{0, -1}, t.GetPlayer())...)
	return pos
}

func (_ *rook) Icon() rune {
	return 'R'
}

func (t *rook) String() string {
	return "rook"
}

func Rook(p core.Player) Token {
	return &rook{protoToken{p, false}}
}

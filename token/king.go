package token

import (
	"github.com/cragcraig/chess/core"
)

type king struct {
	protoToken
}

func (t *king) GetMoves(curPos core.Position, b *Board) []core.Position {
	pos := []core.Position{
		curPos.Add(core.Offset{0, 1}),
		curPos.Add(core.Offset{1, 1}),
		curPos.Add(core.Offset{-1, 1}),
		curPos.Add(core.Offset{0, -1}),
		curPos.Add(core.Offset{1, -1}),
		curPos.Add(core.Offset{-1, -1}),
		curPos.Add(core.Offset{1, 0}),
		curPos.Add(core.Offset{-1, 0}),
	}
	return b.trimInvalidPos(pos, t.GetPlayer())
}

func King(p core.Player) Token {
	return &king{protoToken{p, 0, 'K', "king", -1}}
}

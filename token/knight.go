package token

import (
    "github.com/gnarlyskier/chess/core"
)

type knight struct {
	protoToken
}

func (t *knight) GetMoves(curPos core.Position, b Board) []core.Position {
    pos := []core.Position{
        curPos.Add(core.Offset{1, 2}),
        curPos.Add(core.Offset{2, 1}),
        curPos.Add(core.Offset{1, -2}),
        curPos.Add(core.Offset{2, -1}),
        curPos.Add(core.Offset{-1, 2}),
        curPos.Add(core.Offset{-2, 1}),
        curPos.Add(core.Offset{-1, -2}),
        curPos.Add(core.Offset{-2, -1}),
    }
    return b.trimInvalidPos(pos, t.GetPlayer())
}

func (_ *knight) Icon() rune {
	return 'N'
}

func (t *knight) String() string {
	return string(GetBoardIcon(t))
}

func Knight(p core.Player) Token {
	return &knight{protoToken{p}}
}

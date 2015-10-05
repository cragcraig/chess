package token

import (
    "github.com/gnarlyskier/chess/core"
)

type pawn struct {
	protoToken
}

func (t *pawn) GetMoves(curPos core.Position, b Board) []core.Position {
    pos := []core.Position{}
    // forward move
    forward := curPos.Add(core.Offset{0, t.GetPlayer().ForwardDir()})
    if (b.OnBoard(forward) && b.GetPos(forward) == nil) {
        pos = append(pos, forward)
    }
    // first move
    if !t.HasMoved() {
        forward2 := forward.Add(core.Offset{0, t.GetPlayer().ForwardDir()})
        if (b.OnBoard(forward2) && b.GetPos(forward) == nil && b.GetPos(forward2) == nil) {
            pos = append(pos, forward2)
        }
    }
    // captures
    cap1 := curPos.Add(core.Offset{1, t.GetPlayer().ForwardDir()})
    if (b.isCapture(cap1, t.GetPlayer())) {
        pos = append(pos, cap1)
    }
    cap2 := curPos.Add(core.Offset{-1, t.GetPlayer().ForwardDir()})
    if (b.isCapture(cap2, t.GetPlayer())) {
        pos = append(pos, cap2)
    }
    return pos
}

func (_ *pawn) Icon() rune {
	return 'P'
}

func (t *pawn) String() string {
	return "pawn"
}

func Pawn(p core.Player) Token {
    return &pawn{protoToken{p, false}}
}

package token

import (
    "github.com/gnarlyskier/chess/core"
)

type pawn struct {
	protoToken
}

func (_ *pawn) Icon() rune {
	return 'P'
}

func (t *pawn) String() string {
	return string(GetBoardIcon(t))
}

func Pawn(p core.Player) Token {
	return &pawn{protoToken{p}}
}

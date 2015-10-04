package token

import (
    "github.com/gnarlyskier/chess/core"
)

type rook struct {
	protoToken
}

func (_ *rook) Icon() rune {
	return 'R'
}

func (t *rook) String() string {
	return string(GetBoardIcon(t))
}

func Rook(p core.Player) Token {
	return &rook{protoToken{p}}
}

package token

import (
    "github.com/gnarlyskier/chess/core"
)

type queen struct {
	protoToken
}

func (_ *queen) Icon() rune {
	return 'Q'
}

func (t *queen) String() string {
	return string(GetBoardIcon(t))
}

func Queen(p core.Player) Token {
	return &queen{protoToken{p}}
}

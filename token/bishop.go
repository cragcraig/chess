package token

import (
    "github.com/gnarlyskier/chess/core"
)

type bishop struct {
	protoToken
}

func (_ *bishop) Icon() rune {
	return 'B'
}

func (t *bishop) String() string {
	return string(GetBoardIcon(t))
}

func Bishop(p core.Player) Token {
	return &bishop{protoToken{p}}
}

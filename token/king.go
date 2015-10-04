package token

import (
    "github.com/gnarlyskier/chess/core"
)

type king struct {
	protoToken
}

func (_ *king) Icon() rune {
	return 'K'
}

func (t *king) String() string {
	return string(GetBoardIcon(t))
}

func King(p core.Player) Token {
	return &king{protoToken{p}}
}

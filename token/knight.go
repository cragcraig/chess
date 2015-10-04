package token

import (
    "github.com/gnarlyskier/chess/core"
)

type knight struct {
	protoToken
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

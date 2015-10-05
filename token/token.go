package token

import (
    "github.com/gnarlyskier/chess/core"
)

type Token interface {
	GetPlayer() core.Player
    GetMoves(curPos core.Position, b Board) []core.Position
	String() string
	Icon() rune
}

func GetBoardIcon(t Token) rune {
	return t.GetPlayer().RenderIcon(t.Icon())
}

type protoToken struct {
	player core.Player
}

func (t *protoToken) GetPlayer() core.Player {
	return t.player
}

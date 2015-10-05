package token

import (
	"github.com/gnarlyskier/chess/core"
)

type Token interface {
	GetPlayer() core.Player
	GetMoves(curPos core.Position, b Board) []core.Position
	SetMoved()
	HasMoved() bool
	String() string
	Icon() rune
}

func GetBoardIcon(t Token) rune {
	return t.GetPlayer().RenderIcon(t.Icon())
}

type protoToken struct {
	player core.Player
	moved  bool
}

func (t *protoToken) GetPlayer() core.Player {
	return t.player
}

func (t *protoToken) SetMoved() {
	t.moved = true
}

func (t *protoToken) HasMoved() bool {
	return t.moved
}

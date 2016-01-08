package token

import (
	"github.com/gnarlyskier/chess/core"
)

type Token interface {
	GetPlayer() core.Player
	GetMoves(curPos core.Position, b *Board) []core.Position
	IncMoveCount(currentTurn int)
	HasMoved() bool
	String() string
	Icon() rune
}

func GetBoardIcon(t Token) rune {
	return t.GetPlayer().RenderIcon(t.Icon())
}

type protoToken struct {
	player        core.Player
	moves         int
	icon          rune
	name          string
	lastTurnMoved int
}

func (t *protoToken) GetPlayer() core.Player {
	return t.player
}

func (t *protoToken) IncMoveCount(currentTurn int) {
	t.moves++
	t.lastTurnMoved = currentTurn
}

func (t *protoToken) HasMoved() bool {
	return t.moves != 0
}

func (t *protoToken) Icon() rune {
	return t.icon
}

func (t *protoToken) String() string {
	return t.name
}

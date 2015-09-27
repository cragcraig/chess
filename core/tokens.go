package core

import (
	"strings"
)

type Token interface {
	GetPlayer() Player
	String() string
	Icon() rune
}

func toString(t Token) string {
	if t.GetPlayer() == BLACK {
		return strings.ToUpper(string(t.Icon()))
	}
	return strings.ToLower(string(t.Icon()))
}

type protoToken struct {
	player Player
}

func (t *protoToken) GetPlayer() Player {
	return t.player
}

type King struct {
	protoToken
}

func (_ *King) Icon() rune {
	return 'K'
}

func (t *King) String() string {
	return toString(t)
}

func CreateKing(p Player) Token {
	return &King{protoToken{p}}
}

type Queen struct {
	protoToken
}

func (_ *Queen) Icon() rune {
	return 'Q'
}

func (t *Queen) String() string {
	return toString(t)
}

func CreateQueen(p Player) Token {
	return &Queen{protoToken{p}}
}

type Rook struct {
	protoToken
}

func (_ *Rook) Icon() rune {
	return 'R'
}

func (t *Rook) String() string {
	return toString(t)
}

func CreateRook(p Player) Token {
	return &Rook{protoToken{p}}
}

type Bishop struct {
	protoToken
}

func (_ *Bishop) Icon() rune {
	return 'B'
}

func (t *Bishop) String() string {
	return toString(t)
}

func CreateBishop(p Player) Token {
	return &Bishop{protoToken{p}}
}

type Knight struct {
	protoToken
}

func (_ *Knight) Icon() rune {
	return 'N'
}

func (t *Knight) String() string {
	return toString(t)
}

func CreateKnight(p Player) Token {
	return &Knight{protoToken{p}}
}

type Pawn struct {
	protoToken
}

func (_ *Pawn) Icon() rune {
	return 'P'
}

func (t *Pawn) String() string {
	return toString(t)
}

func CreatePawn(p Player) Token {
	return &Pawn{protoToken{p}}
}

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

type ProtoToken struct {
	player Player
}

func (t *ProtoToken) GetPlayer() Player {
	return t.player
}

type King struct {
	ProtoToken
}

func (_ *King) Icon() rune {
	return 'K'
}

func (t *King) String() string {
	return toString(t)
}

func CreateKing(p Player) Token {
	return &King{ProtoToken{p}}
}

type Queen struct {
	ProtoToken
}

func (_ *Queen) Icon() rune {
	return 'Q'
}

func (t *Queen) String() string {
	return toString(t)
}

func CreateQueen(p Player) Token {
	return &Queen{ProtoToken{p}}
}

type Rook struct {
	ProtoToken
}

func (_ *Rook) Icon() rune {
	return 'R'
}

func (t *Rook) String() string {
	return toString(t)
}

func CreateRook(p Player) Token {
	return &Rook{ProtoToken{p}}
}

type Bishop struct {
	ProtoToken
}

func (_ *Bishop) Icon() rune {
	return 'B'
}

func (t *Bishop) String() string {
	return toString(t)
}

func CreateBishop(p Player) Token {
	return &Bishop{ProtoToken{p}}
}

type Knight struct {
	ProtoToken
}

func (_ *Knight) Icon() rune {
	return 'N'
}

func (t *Knight) String() string {
	return toString(t)
}

func CreateKnight(p Player) Token {
	return &Knight{ProtoToken{p}}
}

type Pawn struct {
	ProtoToken
}

func (_ *Pawn) Icon() rune {
	return 'P'
}

func (t *Pawn) String() string {
	return toString(t)
}

func CreatePawn(p Player) Token {
	return &Pawn{ProtoToken{p}}
}

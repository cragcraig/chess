package core

import (
	"unicode"
)

type Player interface {
	renderIcon(icon rune) rune
}

type white string
type black string

var (
	// Players
	WHITE white = "white"
	BLACK black = "black"
)

func (_ white) renderIcon(icon rune) rune {
	return unicode.ToLower(icon)
}

func (_ black) renderIcon(icon rune) rune {
	return unicode.ToUpper(icon)
}

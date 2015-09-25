package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type TokenType rune

func (t TokenType) String() string {
	return string(rune(t))
}

type RowPos int

func (r RowPos) String() string {
	return string(int(r) + 1)
}

type ColPos int

func (c ColPos) String() string {
	return string(rune(int(c) + int('a')))
}

type Player string

var (
	// Token types
	KING   TokenType = 'K'
	QUEEN  TokenType = 'Q'
	KNIGHT TokenType = 'N'
	ROOK   TokenType = 'R'
	BISHOP TokenType = 'B'
	PAWN   TokenType = 'P'

	// Players
	WHITE Player = "white"
	BLACK Player = "black"

	// Standard starting board
	STANDARD_BOARD Board = Board{
		Width:  8,
		Height: 8,
		Tokens: []Token{
			Rook(BLACK, Pos("a8")),
			Knight(BLACK, Pos("b8")),
			Bishop(BLACK, Pos("c8")),
			Queen(BLACK, Pos("d8")),
			King(BLACK, Pos("e8")),
			Bishop(BLACK, Pos("f8")),
			Knight(BLACK, Pos("g8")),
			Rook(BLACK, Pos("h8")),
		},
	}
)

type Vect struct {
	Col ColPos
	Row RowPos
}

// Converts a position in algebraic notation (e.g., "c4" or "a8") to a Vect
func Pos(pos string) Vect {
	row, err := strconv.Atoi(pos[1:])
	if len(pos) != 2 || !unicode.IsLower(rune(pos[0])) || !unicode.IsDigit(rune(pos[1])) || err != nil {
		panic(fmt.Sprintf("%v is an invalid algebraic notation position", pos))
	}

	return Vect{ColPos(int(pos[0]) - int('a')), RowPos(row - 1)}
}

// Prints in algebraic notation
func (v Vect) String() string {
	return fmt.Sprintf("%v%d", v.Col.String(), v.Row)
}

func (v Vect) Equals(o Vect) bool {
	return v.Col == o.Col && v.Row == o.Row
}

type Token struct {
	Type   TokenType
	Player Player
	Pos    Vect
}

func King(player Player, pos Vect) Token {
	return Token{KING, player, pos}
}

func Queen(player Player, pos Vect) Token {
	return Token{QUEEN, player, pos}
}

func Knight(player Player, pos Vect) Token {
	return Token{KNIGHT, player, pos}
}

func Rook(player Player, pos Vect) Token {
	return Token{ROOK, player, pos}
}

func Bishop(player Player, pos Vect) Token {
	return Token{BISHOP, player, pos}
}

func Pawn(player Player, pos Vect) Token {
	return Token{PAWN, player, pos}
}

type Move struct {
	Orig, Final Vect
}

type Game struct {
	Moves []Move
	Start Board
}

func (g *Game) Add(m Move) {
	g.Moves = append(g.Moves, m)
}

type Board struct {
	Width, Height int
	Tokens        []Token
}

func (b Board) get(v Vect) *Token {
	for i := range b.Tokens {
		if b.Tokens[i].Pos.Equals(v) {
			return &b.Tokens[i]
		}
	}
	return nil
}

func (b Board) String() string {
	out := ""
	for r := b.Height - 1; r >= 0; r-- {
		for c := 0; c < b.Width; c++ {
			if t := b.get(Vect{ColPos(c), RowPos(r)}); t != nil {
				out += t.Type.String()
			} else {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}

func main() {
	fmt.Println(STANDARD_BOARD)
}

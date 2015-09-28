package core

import (
	"fmt"
	"strconv"
	"unicode"
)

type RowPos int

func (r RowPos) GetBoardHeading() string {
	return strconv.Itoa(int(r) + 1)
}

func (r RowPos) String() string {
	return string(int(r))
}

func (r RowPos) Next() RowPos {
	return RowPos(int(r) - 1)
}

type ColPos int

func (c ColPos) GetBoardHeading() rune {
	return rune(int(c) + int('a'))
}

func (r ColPos) String() string {
	return string(int(r))
}

func (r ColPos) Next() ColPos {
	return ColPos(int(r) + 1)
}

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

func (v Vect) String() string {
	return fmt.Sprintf("(%s, %s)", v.Col, v.Row)
}

// Prints in algebraic notation
func (v Vect) AsPos() string {
	return fmt.Sprintf("%s%s", v.Col.GetBoardHeading(), v.Row.GetBoardHeading())
}

func (v Vect) Equals(o Vect) bool {
	return v.Col == o.Col && v.Row == o.Row
}

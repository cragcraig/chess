package core

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

type Row int

func (r Row) GetBoardHeading() string {
	return strconv.Itoa(int(r) + 1)
}

func (r Row) String() string {
	return string(int(r))
}

func (r Row) Next() Row {
	return Row(int(r) - 1)
}

type Column int

func (c Column) GetBoardHeading() rune {
	return rune(int(c) + int('a'))
}

func (r Column) String() string {
	return string(int(r))
}

func (r Column) Next() Column {
	return Column(int(r) + 1)
}

type Position interface {
	Column() Column
	Row() Row
	Equals(Position) bool
	AsAlgPos() string
}

type position struct {
	col Column
	row Row
}

func FromCoordPos(c Column, r Row) Position {
	return position{c, r}
}

// Converts a position in algebraic notation (e.g., "c4" or "a8") to a Position
func FromAlgPos(pos string) (Position, error) {
	row, err := strconv.Atoi(pos[1:])
	if len(pos) != 2 || !unicode.IsLower(rune(pos[0])) || !unicode.IsDigit(rune(pos[1])) || err != nil {
		return nil, errors.New(fmt.Sprintf("%v is an invalid algebraic notation position", pos))
	}
	return FromCoordPos(Column(int(pos[0])-int('a')), Row(row-1)), nil
}

func (v position) Column() Column {
	return v.col
}

func (v position) Row() Row {
	return v.row
}

func (v position) String() string {
	return fmt.Sprintf("(%s, %s)", v.col, v.row)
}

// Prints in algebraic notation
func (v position) AsAlgPos() string {
	return fmt.Sprintf("%s%s", v.col.GetBoardHeading(), v.row.GetBoardHeading())
}

func (v position) Equals(o Position) bool {
	return v.col == o.Column() && v.row == o.Row()
}

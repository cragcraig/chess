package core

import (
	"fmt"
)

type Board struct {
	width, height int
	tokens        []Token
}

func CreateEmptyBoard(w, h int) Board {
	if w > 26 || h > 99 || w < 0 || h < 0 {
		// No real reason for this limit except we haven't bothered to
		// provide row/column headings for larger boards.
		panic("Boards can not be larger than 26x99")
	}
	return Board{
		width:  w,
		height: h,
		tokens: make([]Token, w*h),
	}
}

func CreateStandardBoard() Board {
	b := CreateEmptyBoard(8, 8)
	// black
	b.put(CreateRook(BLACK), FromCoordPos(0, 7))
	b.put(CreateKnight(BLACK), FromCoordPos(1, 7))
	b.put(CreateBishop(BLACK), FromCoordPos(2, 7))
	b.put(CreateQueen(BLACK), FromCoordPos(3, 7))
	b.put(CreateKing(BLACK), FromCoordPos(4, 7))
	b.put(CreateBishop(BLACK), FromCoordPos(5, 7))
	b.put(CreateKnight(BLACK), FromCoordPos(6, 7))
	b.put(CreateRook(BLACK), FromCoordPos(7, 7))
	for i := 0; i < 8; i++ {
		b.put(CreatePawn(BLACK), FromCoordPos(Column(i), 6))
	}
	// white
	b.put(CreateRook(WHITE), FromCoordPos(0, 0))
	b.put(CreateKnight(WHITE), FromCoordPos(1, 0))
	b.put(CreateBishop(WHITE), FromCoordPos(2, 0))
	b.put(CreateQueen(WHITE), FromCoordPos(3, 0))
	b.put(CreateKing(WHITE), FromCoordPos(4, 0))
	b.put(CreateBishop(WHITE), FromCoordPos(5, 0))
	b.put(CreateKnight(WHITE), FromCoordPos(6, 0))
	b.put(CreateRook(WHITE), FromCoordPos(7, 0))
	for i := 0; i < 8; i++ {
		b.put(CreatePawn(WHITE), FromCoordPos(Column(i), 1))
	}
	return b
}

func (b Board) GetPos(v Position) Token {
	return b.tokens[b.indexPos(v)]
}

func (b Board) Get(c Column, r Row) Token {
	return b.tokens[b.index(c, r)]
}

func (b Board) put(t Token, v Position) {
	b.tokens[b.indexPos(v)] = t
}

func (b Board) indexPos(v Position) int {
	return b.index(v.Column(), v.Row())
}

func (b Board) index(c Column, r Row) int {
	return int(r)*b.width + int(c)
}

func (b Board) FirstCol() Column {
	return Column(0)
}

func (b Board) EndCol() Column {
	return Column(b.width)
}

func (b Board) FirstRow() Row {
	return Row(b.height - 1)
}

func (b Board) EndRow() Row {
	return Row(-1)
}

func (b Board) String() string {
	// column ids
	out := []rune{' ', ' ', ' ', ' '}
	for c := b.FirstCol(); c != b.EndCol(); c = c.Next() {
		out = append(out, ' ')
		out = append(out, c.GetBoardHeading())
		out = append(out, ' ')
	}
	out = append(out, []rune{'\n', '\n'}...)
	for r := b.FirstRow(); r != b.EndRow(); r = r.Next() {
		// row ids
		out = append(out, []rune(fmt.Sprintf("%2s", r.GetBoardHeading()))...)
		out = append(out, []rune{' ', ' '}...)
		for c := b.FirstCol(); c != b.EndCol(); c = c.Next() {
			out = append(out, ' ')
			// icon for this position
			if t := b.Get(c, r); t != nil {
				out = append(out, GetBoardIcon(t))
			} else {
				out = append(out, '.')
			}
			out = append(out, ' ')
		}
		out = append(out, '\n')
	}
	return string(out)
}

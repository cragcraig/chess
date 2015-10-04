package game

import (
	"fmt"
    "github.com/gnarlyskier/chess/core"
    "github.com/gnarlyskier/chess/token"
)

type Board struct {
	width, height int
	tokens        []token.Token
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
		tokens: make([]token.Token, w*h),
	}
}

func CreateStandardBoard() Board {
	b := CreateEmptyBoard(8, 8)
	// black
	b.put(token.Rook(core.BLACK), core.FromCoordPos(0, 7))
	b.put(token.Knight(core.BLACK), core.FromCoordPos(1, 7))
	b.put(token.Bishop(core.BLACK), core.FromCoordPos(2, 7))
	b.put(token.Queen(core.BLACK), core.FromCoordPos(3, 7))
	b.put(token.King(core.BLACK), core.FromCoordPos(4, 7))
	b.put(token.Bishop(core.BLACK), core.FromCoordPos(5, 7))
	b.put(token.Knight(core.BLACK), core.FromCoordPos(6, 7))
	b.put(token.Rook(core.BLACK), core.FromCoordPos(7, 7))
	for i := 0; i < 8; i++ {
		b.put(token.Pawn(core.BLACK), core.FromCoordPos(core.Column(i), 6))
	}
	// white
	b.put(token.Rook(core.WHITE), core.FromCoordPos(0, 0))
	b.put(token.Knight(core.WHITE), core.FromCoordPos(1, 0))
	b.put(token.Bishop(core.WHITE), core.FromCoordPos(2, 0))
	b.put(token.Queen(core.WHITE), core.FromCoordPos(3, 0))
	b.put(token.King(core.WHITE), core.FromCoordPos(4, 0))
	b.put(token.Bishop(core.WHITE), core.FromCoordPos(5, 0))
	b.put(token.Knight(core.WHITE), core.FromCoordPos(6, 0))
	b.put(token.Rook(core.WHITE), core.FromCoordPos(7, 0))
	for i := 0; i < 8; i++ {
		b.put(token.Pawn(core.WHITE), core.FromCoordPos(core.Column(i), 1))
	}
	return b
}

func (b Board) GetPos(v core.Position) token.Token {
	return b.tokens[b.indexPos(v)]
}

func (b Board) Get(c core.Column, r core.Row) token.Token {
	return b.tokens[b.index(c, r)]
}

func (b Board) put(t token.Token, v core.Position) {
	b.tokens[b.indexPos(v)] = t
}

func (b Board) indexPos(v core.Position) int {
	return b.index(v.Column(), v.Row())
}

func (b Board) index(c core.Column, r core.Row) int {
	return int(r)*b.width + int(c)
}

func (b Board) OnBoard(v core.Position) bool {
	c := int(v.Column())
	r := int(v.Row())
	return c >= 0 && c < b.width && r >= 0 && r < b.height
}

func (b Board) FirstCol() core.Column {
	return core.Column(0)
}

func (b Board) EndCol() core.Column {
	return core.Column(b.width)
}

func (b Board) FirstRow() core.Row {
	return core.Row(b.height - 1)
}

func (b Board) EndRow() core.Row {
	return core.Row(-1)
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
				out = append(out, token.GetBoardIcon(t))
			} else {
				out = append(out, '.')
			}
			out = append(out, ' ')
		}
		out = append(out, '\n')
	}
	return string(out)
}

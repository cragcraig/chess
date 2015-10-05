package token

import (
	"fmt"
	"github.com/gnarlyskier/chess/core"
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
	b.putOrig(Rook(core.BLACK), core.CoordPos(0, 7))
	b.putOrig(Knight(core.BLACK), core.CoordPos(1, 7))
	b.putOrig(Bishop(core.BLACK), core.CoordPos(2, 7))
	b.putOrig(Queen(core.BLACK), core.CoordPos(3, 7))
	b.putOrig(King(core.BLACK), core.CoordPos(4, 7))
	b.putOrig(Bishop(core.BLACK), core.CoordPos(5, 7))
	b.putOrig(Knight(core.BLACK), core.CoordPos(6, 7))
	b.putOrig(Rook(core.BLACK), core.CoordPos(7, 7))
	for i := 0; i < 8; i++ {
		b.putOrig(Pawn(core.BLACK), core.CoordPos(core.Column(i), 6))
	}
	// white
	b.putOrig(Rook(core.WHITE), core.CoordPos(0, 0))
	b.putOrig(Knight(core.WHITE), core.CoordPos(1, 0))
	b.putOrig(Bishop(core.WHITE), core.CoordPos(2, 0))
	b.putOrig(Queen(core.WHITE), core.CoordPos(3, 0))
	b.putOrig(King(core.WHITE), core.CoordPos(4, 0))
	b.putOrig(Bishop(core.WHITE), core.CoordPos(5, 0))
	b.putOrig(Knight(core.WHITE), core.CoordPos(6, 0))
	b.putOrig(Rook(core.WHITE), core.CoordPos(7, 0))
	for i := 0; i < 8; i++ {
		b.putOrig(Pawn(core.WHITE), core.CoordPos(core.Column(i), 1))
	}
	return b
}

func (b Board) GetPos(v core.Position) Token {
	return b.tokens[b.indexPos(v)]
}

func (b Board) Get(c core.Column, r core.Row) Token {
	return b.tokens[b.index(c, r)]
}

func (b Board) putOrig(t Token, v core.Position) {
	b.tokens[b.indexPos(v)] = t
}

func (b Board) put(t Token, v core.Position) {
	b.putOrig(t, v)
	if t != nil {
		t.SetMoved()
	}
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

func (b Board) DoMove(move core.Move) {
	t := b.GetPos(move.Orig)
	b.put(t, move.Final)
	b.put(nil, move.Orig)
}

func (b Board) IsValidMove(move core.Move, player core.Player) bool {
	if t := b.GetPos(move.Orig); t != nil && t.GetPlayer() == player {
		moves := t.GetMoves(move.Orig, b)
		for i := range moves {
			if move.Final.Equals(moves[i]) {
				return true
			}
		}
	}
	return false
}

func (b Board) validPos(pos core.Position, player core.Player) bool {
	if b.OnBoard(pos) {
		if t := b.GetPos(pos); t == nil || t.GetPlayer() != player {
			return true
		}
	}
	return false
}

func (b Board) isCapture(pos core.Position, player core.Player) bool {
	if b.OnBoard(pos) {
		if t := b.GetPos(pos); t != nil && t.GetPlayer() != player {
			return true
		}
	}
	return false
}

func (b Board) trimInvalidPos(pos []core.Position, player core.Player) []core.Position {
	filtered := []core.Position{}
	for i := range pos {
		if b.validPos(pos[i], player) {
			filtered = append(filtered, pos[i])
		}
	}
	return filtered
}

func (b Board) extendMove(pos core.Position, offset core.Offset, player core.Player) []core.Position {
	newPos := []core.Position{}
	cur := pos
	for {
		cur = cur.Add(offset)
		if !b.validPos(cur, player) {
			break
		}
		newPos = append(newPos, cur)
	}
	return newPos
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
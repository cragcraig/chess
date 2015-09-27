package core

type Board struct {
	width, height int
	tokens        []Token
}

func CreateEmptyBoard(w, h int) Board {
	return Board{
		width:  w,
		height: h,
		tokens: make([]Token, w*h),
	}
}

func CreateStandardBoard() Board {
	b := CreateEmptyBoard(8, 8)
	// black
	b.put(CreateRook(BLACK), Pos("a8"))
	b.put(CreateKnight(BLACK), Pos("b8"))
	b.put(CreateBishop(BLACK), Pos("c8"))
	b.put(CreateQueen(BLACK), Pos("d8"))
	b.put(CreateKing(BLACK), Pos("e8"))
	b.put(CreateBishop(BLACK), Pos("f8"))
	b.put(CreateKnight(BLACK), Pos("g8"))
	b.put(CreateRook(BLACK), Pos("h8"))
	for i := 0; i < 8; i++ {
		b.put(CreatePawn(BLACK), Vect{ColPos(i), 6})
	}
	// white
	b.put(CreateRook(WHITE), Pos("a1"))
	b.put(CreateKnight(WHITE), Pos("b1"))
	b.put(CreateBishop(WHITE), Pos("c1"))
	b.put(CreateQueen(WHITE), Pos("d1"))
	b.put(CreateKing(WHITE), Pos("e1"))
	b.put(CreateBishop(WHITE), Pos("f1"))
	b.put(CreateKnight(WHITE), Pos("g1"))
	b.put(CreateRook(WHITE), Pos("h1"))
	for i := 0; i < 8; i++ {
		b.put(CreatePawn(WHITE), Vect{ColPos(i), 1})
	}
	return b
}

func (b Board) Get(v Vect) Token {
	return b.tokens[b.index(v)]
}

func (b Board) put(t Token, v Vect) {
	b.tokens[b.index(v)] = t
}

func (b Board) index(v Vect) int {
	return int(v.Row)*b.width + int(v.Col)
}

func (b Board) FirstCol() ColPos {
	return ColPos(0)
}

func (b Board) EndCol() ColPos {
	return ColPos(b.width)
}

func (b Board) FirstRow() RowPos {
	return RowPos(b.height - 1)
}

func (b Board) EndRow() RowPos {
	return RowPos(-1)
}

func (b Board) String() string {
	// column ids
	out := []rune{' ', ' ', ' ', ' '}
	for c := b.FirstCol(); c != b.EndCol(); c = c.Next() {
		out = append(out, ' ')
		out = append(out, []rune(c.String())...)
		out = append(out, ' ')
	}
	out = append(out, []rune{'\n', '\n'}...)
	for r := b.FirstRow(); r != b.EndRow(); r = r.Next() {
		// row ids
		out = append(out, ' ')
		out = append(out, []rune(r.String())...)
		out = append(out, []rune{' ', ' '}...)
		for c := b.FirstCol(); c != b.EndCol(); c = c.Next() {
			out = append(out, ' ')
			// icon for this position
			if t := b.Get(Vect{c, r}); t != nil {
				out = append(out, []rune(t.String())...)
			} else {
				out = append(out, '.')
			}
			out = append(out, ' ')
		}
		out = append(out, '\n')
	}
	return string(out)
}

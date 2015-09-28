package core

type Token interface {
	GetPlayer() Player
	String() string
	Icon() rune
}

func GetBoardIcon(t Token) rune {
	return t.GetPlayer().renderIcon(t.Icon())
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
	return string(GetBoardIcon(t))
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
	return string(GetBoardIcon(t))
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
	return string(GetBoardIcon(t))
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
	return string(GetBoardIcon(t))
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
	return string(GetBoardIcon(t))
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
	return string(GetBoardIcon(t))
}

func CreatePawn(p Player) Token {
	return &Pawn{protoToken{p}}
}

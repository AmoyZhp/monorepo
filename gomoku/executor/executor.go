package executor

// Player game player
type Player int

// player num
const (
	EMPTY Player = 0
	BLACK Player = 1
	WHITE Player = 2
)

// Move chess act
type Move struct {
	Col    int    `json:"col"`
	Row    int    `json:"row"`
	Player Player `json:"player"`
}

// Executor compose other elements
type Executor interface {
	GetNextMove(Move) Move
}

// Gomoku executor implement
type Gomoku struct {
}

// GetNextMove get next move
func (imp *Gomoku) GetNextMove(move Move) Move {
	return Move{}
}

// NewGomokuExectuor new gomoku exectuor
func NewGomokuExectuor() Executor {
	return &Gomoku{}
}

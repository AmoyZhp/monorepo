package main

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

// Engine algorithm engine
type Engine interface {
	Predict(Move) *Move
}

// Board chessboard
type Board struct {
	board       [][]Player
	historyMove []Move
	timeline    int
}

// Set set chess on board
func (imp *Board) Set(x, y int, player Player) error {
	// TODO imp set

	return nil
}

// Regret move back one step
func (imp *Board) Regret() error {
	// TODO imp Regret.
	return nil
}

// Reset reset borad
func (imp *Board) Reset() error {
	// TODO imp rest.
	return nil
}

// TreeSearchEngine implement tree serch engine
type TreeSearchEngine struct {
	board Board
}

// Predict get next move
func (imp *TreeSearchEngine) Predict(move Move) *Move {
	// TODO implement predict logic
	err := imp.board.Set(move.Col, move.Row, move.Player)
	if err != nil {
		return nil
	}
	return imp.search(imp.board, imp.nextPlayer(move.Player))
}

// search search next move
func (imp *TreeSearchEngine) search(board Board, player Player) *Move {
	// TODO imp search
	return nil
}

func (imp *TreeSearchEngine) nextPlayer(player Player) Player {
	if player == WHITE {
		return BLACK
	}
	return WHITE
}

func newEngine() Engine {
	return &TreeSearchEngine{}
}

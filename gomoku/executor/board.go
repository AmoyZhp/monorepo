package executor

import "fmt"

// Board chessboard
type Board interface {
	Set(Move) error
	GetPlayerAtPos(x, y int) (Player, error)
	Regret(step int) error
	Reset() error
}

// GomokuBoard chessboard
type GomokuBoard struct {
	board       [][]Player
	historyMove []Move
	timeline    int
	maxCol      int
	maxRow      int
}

// Set set chess on board
func (imp *GomokuBoard) Set(move Move) error {
	// TODO imp set

	return nil
}

// Regret move back n step
func (imp *GomokuBoard) Regret(n int) error {
	// TODO imp Regret.
	return nil
}

// Reset reset borad
func (imp *GomokuBoard) Reset() error {
	// TODO imp rest.
	return nil
}

// GetPlayerAtPos get player at specific position
func GetPlayerAtPos(x, y int) (Player, error) {
	return EMPTY, fmt.Errorf("")
}

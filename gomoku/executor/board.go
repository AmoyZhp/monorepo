package executor

import (
	"fmt"
)

// Board chessboard
type Board interface {
	Set(Move) error
	GetPlayerAtPos(x, y int) (Player, error)
	Regret(step int) error
	Reset() error
	IsEnd() bool
}

// Gomoku related attribute
const (
	GomokuCol = 15
	GomokuRow = 15
)

// NewGomokuBoard new gomoku chess board
func NewGomokuBoard() Board {
	borad := make([][]Player, GomokuRow)
	for i := 0; i < GomokuRow; i++ {
		borad[i] = make([]Player, GomokuCol)
	}
	historyMove := make([]Move, GomokuCol*GomokuRow)
	gomokuBoard := &GomokuBoard{
		board:       borad,
		historyMove: historyMove,
		maxCol:      GomokuCol,
		maxRow:      GomokuRow,
		timeline:    0,
		endTime:     GomokuCol * GomokuRow,
	}
	gomokuBoard.Reset()
	return gomokuBoard
}

// GomokuBoard chessboard
type GomokuBoard struct {
	board       [][]Player
	historyMove []Move
	timeline    int
	endTime     int
	maxCol      int
	maxRow      int
}

// Set set chess on board
func (imp *GomokuBoard) Set(move Move) error {
	if move.Player != BLACK && move.Player != WHITE {
		return fmt.Errorf("playe is invalid. get player: %d", move.Player)
	}
	if move.Col < 0 || move.Col > imp.maxCol {
		return fmt.Errorf("column position is invalid. get column: %d", move.Col)
	}
	if move.Row < 0 || move.Row > imp.maxRow {
		return fmt.Errorf("row position is invalid. get row: %d", move.Row)
	}
	if imp.board[move.Row][move.Col] != EMPTY {
		return fmt.Errorf("(%d, %d) already taken by player %d", move.Row, move.Col, imp.board[move.Row][move.Col])
	}

	imp.board[move.Row][move.Col] = move.Player

	return nil
}

// Regret move back n step
func (imp *GomokuBoard) Regret(n int) error {
	// TODO imp Regret.
	return nil
}

// Reset reset borad
func (imp *GomokuBoard) Reset() error {
	for i := 0; i < imp.maxCol; i++ {
		for j := 0; j < imp.maxRow; j++ {
			imp.board[i][j] = EMPTY
		}
	}
	imp.historyMove = make([]Move, imp.maxCol*imp.maxRow)
	imp.timeline = 0
	return nil
}

// IsEnd the game is over or not
func (imp *GomokuBoard) IsEnd() bool {
	if imp.timeline == imp.endTime {
		return true
	}
	return false
}

// GetPlayerAtPos get player at specific position
func (imp *GomokuBoard) GetPlayerAtPos(row, col int) (Player, error) {
	if col < 0 || col > imp.maxCol {
		return EMPTY, fmt.Errorf("column position is invalid. get column: %d", col)
	}
	if row < 0 || row > imp.maxRow {
		return EMPTY, fmt.Errorf("row position is invalid. get row: %d", row)
	}

	return imp.board[row][col], nil
}

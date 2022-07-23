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
	Eval() int
	GetPosInFour(Player) []Move
	GetPosInLiveThree(Player) []Move
	GetPosInThree(Player) []Move
	GetPosInTwo(Player) []Move
	GetPosInOne(Player) []Move
	GetOpenForm(Player) []Move
}

// Gomoku related attribute
const (
	GomokuCol = 15
	GomokuRow = 15
)

type direction []int

// Row return row axis direction
func (imp direction) Row() int {
	return imp[0]
}

// Col return column axis direction
func (imp direction) Col() int {
	return imp[1]
}

var (
	northToSouth         = direction{1, 0}
	westToEast           = direction{0, 1}
	northWestToSouthEast = direction{1, 1}
	northEastToSouthWest = direction{1, -1}
)

var directions = []direction{
	northToSouth,
	westToEast,
	northWestToSouthEast,
	northEastToSouthWest,
}

// NewGomokuBoard new gomoku chess board
func NewGomokuBoard() Board {
	return newGomokuBoard()
}

func newGomokuBoard() *GomokuBoard {
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
	roardBoard  RoadBoard
}

// Set set chess on board
func (imp *GomokuBoard) Set(move Move) error {
	if err := imp.validCoordinate(move.Row, move.Col); err != nil {
		return err
	}
	if move.Player != BLACK && move.Player != WHITE {
		return fmt.Errorf("player is invalid. get player: %d", move.Player)
	}
	if imp.board[move.Row][move.Col] != EMPTY {
		return fmt.Errorf("(%d, %d) already taken by player %d", move.Row, move.Col, imp.board[move.Row][move.Col])
	}

	imp.board[move.Row][move.Col] = move.Player
	err := imp.roardBoard.Set(move)
	if err != nil {
		fmt.Println("roadboard update error: ", err.Error())
	}
	imp.timeline++
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
	if err := imp.validCoordinate(row, col); err != nil {
		return EMPTY, err
	}
	return imp.board[row][col], nil
}

// Eval evalutate the borad status
func (imp *GomokuBoard) Eval() int {
	return -1
}

func (imp *GomokuBoard) validCoordinate(row, col int) error {
	if col < 0 || col > imp.maxCol {
		return fmt.Errorf("column position is invalid. get column: %d", col)
	}
	if row < 0 || row > imp.maxRow {
		return fmt.Errorf("row position is invalid. get row: %d", row)
	}
	return nil
}

func (imp *GomokuBoard) countNewRoad(move Move) []int {
	roadCount := make([]int, 6)
	for _, d := range directions {
		startRow := move.Row - 4*d.Row()
		startCol := move.Col - 4*d.Col()
		for i := 0; i < 5; i++ {
			roadNum := imp.computeRoadNum(startRow+i*d.Row(), startCol+i*d.Col(), move.Player, d)
			roadCount[roadNum]++
		}
	}
	return roadCount
}

func (imp *GomokuBoard) computeRoadNum(startRow, startCol int, player Player, d direction) int {
	playerCnt := 0
	for i := 0; i < 5; i++ {
		row := startRow + i*d.Row()
		col := startCol + i*d.Col()
		if imp.validCoordinate(row, col) != nil ||
			imp.board[row][col] == nextPlayer(player) {
			return 0
		}
		if imp.board[row][col] == player {
			playerCnt++
		}
	}
	return playerCnt
}

// GetPosInFour get moves in four road
func (imp *GomokuBoard) GetPosInFour(player Player) []Move {
	return imp.roardBoard.GetPosInFour(player)
}

// GetPosInLiveThree get moves in live three
func (imp *GomokuBoard) GetPosInLiveThree(player Player) []Move {
	return imp.roardBoard.GetPosInLiveThree(player)
}

// GetPosInThree get moves in three road
func (imp *GomokuBoard) GetPosInThree(player Player) []Move {
	return imp.roardBoard.GetPosInThree(player)
}

// GetPosInTwo get moves in two road
func (imp *GomokuBoard) GetPosInTwo(player Player) []Move {
	return imp.roardBoard.GetPosInTwo(player)
}

// GetPosInOne get moves in one road
func (imp *GomokuBoard) GetPosInOne(player Player) []Move {
	return imp.roardBoard.GetPosInOne(player)
}

// GetOpenForm get open form
func (imp *GomokuBoard) GetOpenForm(_ Player) []Move {
	panic("not implemented") // TODO: Implement
}

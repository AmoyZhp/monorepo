package executor

// Engine algorithm engine
type Engine interface {
	Predict(Move) *Move
}

// TreeSearchEngine implement tree serch engine
type TreeSearchEngine struct {
	board Board
}

// Predict get next move
func (imp *TreeSearchEngine) Predict(move Move) *Move {
	// TODO implement predict logic
	err := imp.board.Set(move)
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

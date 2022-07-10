package executor

import (
	"fmt"
)

// Engine algorithm engine
type Engine interface {
	Predict(Move) (*Move, error)
}

// Evaluator to evaluate a game status
type Evaluator interface {
	Evaluate(Board) int
}

// TreeSearchEngine implement tree serch engine
type TreeSearchEngine struct {
	board     Board
	evaluator Evaluator
}

// Predict get next move
func (imp *TreeSearchEngine) Predict(move Move) (*Move, error) {
	// TODO implement predict logic
	err := imp.board.Set(move)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return imp.search(imp.board, nextPlayer(move.Player))
}

// search search next move
func (imp *TreeSearchEngine) search(board Board, player Player) (*Move, error) {
	// TODO imp search
	return nil, fmt.Errorf("")
}

func (imp *TreeSearchEngine) minmax(
	board Board,
	masterPlayer, actingPlayer Player,
	depth, alpha, beta int,
) int {
	if depth == 0 || board.IsEnd() {
		return imp.evaluator.Evaluate(board)
	}
	moves := imp.findBestMoves(board, actingPlayer)
	for _, m := range moves {
		board.Set(m)
		eval := imp.minmax(board, masterPlayer, nextPlayer(actingPlayer), depth-1, alpha, beta)
		board.Regret(1)
		alpha, beta = upadteAlphaBeta(masterPlayer, actingPlayer, eval, alpha, beta)
		if alpha > beta {
			break
		}
	}
	if isMaxLevel(masterPlayer, actingPlayer) {
		return alpha
	}
	return beta
}
func isMaxLevel(masterPlayer, actingPlayer Player) bool {
	if masterPlayer == actingPlayer {
		return true
	}
	return false
}

func upadteAlphaBeta(
	masterPlayer, actingPlayer Player,
	eval, alpha, beta int,
) (int, int) {
	if isMaxLevel(masterPlayer, actingPlayer) {
		if eval > alpha {
			alpha = eval
		}
	} else {
		if eval < beta {
			beta = eval
		}
	}
	return alpha, beta
}

func (imp *TreeSearchEngine) findBestMoves(board Board, actingPlayer Player) []Move {
	// TODO imp find best moves
	return nil
}

func newEngine() Engine {
	return &TreeSearchEngine{}
}

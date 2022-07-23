package executor

import (
	"fmt"
	"math"
)

// Engine algorithm engine
type Engine interface {
	Predict(Move) (*Move, error)
}

// Evaluator to evaluate a game status
type Evaluator interface {
	Evaluate(Board) int
}

// TreeSearchEngineConfig engine config
type TreeSearchEngineConfig struct {
	searchDepth int
}

// TreeSearchEngine implement tree serch engine
type TreeSearchEngine struct {
	board     Board
	evaluator Evaluator
	conf      TreeSearchEngineConfig
}

// Predict get next move
func (imp *TreeSearchEngine) Predict(move Move) (*Move, error) {
	// to set opponet move
	err := imp.board.Set(move)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return imp.search(imp.board, nextPlayer(move.Player))
}

// search search next move
func (imp *TreeSearchEngine) search(board Board, player Player) (*Move, error) {
	alpha := math.MinInt
	beta := math.MaxInt
	depth := imp.conf.searchDepth
	moves := imp.findBestMoves(board, player)
	if len(moves) == 0 {
		return nil, fmt.Errorf("can not find candidate moves")
	}
	var bestMoves *Move
	for _, m := range moves {
		board.Set(m)
		eval := imp.alphaBetaPruning(board, player, nextPlayer(player), depth-1, alpha, beta)
		board.Regret(1)
		if eval > alpha {
			alpha = eval
			bestMoves = &m
		}
	}
	// bestMoves will not be nil
	// beacuse the first moves value certainly greater than alpha, bestMoves will be sat
	return bestMoves, nil
}

func (imp *TreeSearchEngine) alphaBetaPruning(
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
		eval := imp.alphaBetaPruning(board, masterPlayer, nextPlayer(actingPlayer), depth-1, alpha, beta)
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
	var moves []Move
	if moves = board.GetPosInFour(actingPlayer); moves != nil {
		return moves
	}
	if moves = board.GetPosInFour(nextPlayer(actingPlayer)); moves != nil {
		return moves
	}
	if moves = board.GetPosInLiveThree(actingPlayer); moves != nil {
		return moves
	}
	if moves = board.GetPosInLiveThree(nextPlayer(actingPlayer)); moves != nil {
		return moves
	}
	moves = make([]Move, 0)
	moves = append(moves, board.GetPosInThree(actingPlayer))
	moves = append(moves, board.GetPosInThree(nextPlayer(actingPlayer)))
	moves = append(moves, board.GetPosInTwo(actingPlayer))
	moves = append(moves, board.GetPosInTwo(nextPlayer(actingPlayer)))
	moves = append(moves, board.GetPosInOne(actingPlayer))
	moves = append(moves, board.GetPosInOne(nextPlayer(actingPlayer)))
	if moves != nil {
		return moves
	}
	moves = board.GetOpenForm(actingPlayer)
	return moves
}

func newEngine() Engine {
	return &TreeSearchEngine{}
}

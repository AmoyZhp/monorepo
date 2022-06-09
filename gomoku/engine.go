package main

// Engine algorithm engine
type Engine interface {
	Predict(Move) Move
}

// TreeSearchEngine implement tree serch engine
type TreeSearchEngine struct {
}

// Predict get next move
func (imp TreeSearchEngine) Predict(move Move) Move {
	// TODO implement predict logic
	return Move{Col: move.Col + 1, Row: move.Row + 1}
}

func newEngine() Engine {
	return TreeSearchEngine{}
}

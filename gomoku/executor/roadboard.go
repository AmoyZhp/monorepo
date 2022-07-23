package executor

// RoadBoard board to store roads
type RoadBoard struct{}

// Set update board when board updated
func (imp *RoadBoard) Set(move Move) error {
	panic("unimplement")
}

// Regret move back n step
func (imp *RoadBoard) Regret(step int) error {
	panic("unimplement")
}

// GetPosInFour get positin in four road
func (imp *RoadBoard) GetPosInFour(player Player) []Move {
	panic("unimplement")
}

// GetPosInLiveThree get live three position in three road
func (imp *RoadBoard) GetPosInLiveThree(player Player) []Move {
	panic("unimplement")
}

// GetPosInThree get position in three road
func (imp *RoadBoard) GetPosInThree(player Player) []Move {
	panic("unimplement")
}

// GetPosInTwo get position in two road
func (imp *RoadBoard) GetPosInTwo(player Player) []Move {
	panic("unimplement")
}

// GetPosInOne get position in one road
func (imp *RoadBoard) GetPosInOne(player Player) []Move {
	panic("unimplement")
}

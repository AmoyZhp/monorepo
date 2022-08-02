package executor

// RoadBoard board to store roads
type RoadBoard struct {
	roadsPool      [][]Road
	roads          map[Player][][]Road
	liveThreeRoads map[Player][]Road
}

// Road road is continuous five position in specific direction
type Road struct {
}

// NewRoadBoard new road board
func NewRoadBoard() RoadBoard {

	roadsPool := newRoadsPool()
	roads := newRoads(roadsPool)
	liveThreeRoads := newLiveThreeRoads(roads)
	return RoadBoard{
		roadsPool:      roadsPool,
		roads:          roads,
		liveThreeRoads: liveThreeRoads,
	}
}

func newRoadsPool() [][]Road {
	panic("unimplement")
}

func newRoads(roadsPool [][]Road) map[Player][][]Road {
	panic("unimplement")
}

func newLiveThreeRoads(map[Player][][]Road) map[Player][]Road {
	panic("unimplement")
}

func (imp *Road) getEmptyPos() []Move {
	panic("unimplement")
}

// Set update board when board updated
func (imp *RoadBoard) Set(move Move) error {
	panic("unimplement")
}

// Regret move back n step
func (imp *RoadBoard) Regret(step int) error {
	panic("unimplement")
}

// GetPosInLiveThree get live three position in three road
func (imp *RoadBoard) GetPosInLiveThree(player Player) []Move {
	moves := make([]Move, 0)
	for _, r := range imp.liveThreeRoads[player] {
		moves = append(moves, r.getEmptyPos()...)
	}
	return moves
}

// GetPosInFour get positin in four road
func (imp *RoadBoard) GetPosInFour(player Player) []Move {
	return imp.getPos(player, 4)
}

// GetPosInThree get position in three road
func (imp *RoadBoard) GetPosInThree(player Player) []Move {
	return imp.getPos(player, 3)
}

// GetPosInTwo get position in two road
func (imp *RoadBoard) GetPosInTwo(player Player) []Move {
	return imp.getPos(player, 2)
}

// GetPosInOne get position in one road
func (imp *RoadBoard) GetPosInOne(player Player) []Move {
	return imp.getPos(player, 1)
}

func (imp *RoadBoard) getPos(player Player, roadNum int) []Move {
	moves := make([]Move, 0)
	for _, r := range imp.roads[player][roadNum] {
		moves = append(moves, r.getEmptyPos()...)
	}
	return moves
}

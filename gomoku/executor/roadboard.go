package executor

// RoadsLayer roads layer
type RoadsLayer map[Player][][]*Road

// RoadBoard board to store roads
type RoadBoard struct {
	roadsPool      [][]*Road
	roads          RoadsLayer
	liveThreeRoads map[Player][]*Road
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

func newRoadsPool() [][]*Road {
	panic("unimplement")
}

func newRoads(roadsPool [][]*Road) map[Player][][]*Road {
	panic("unimplement")
}

func newLiveThreeRoads(map[Player][][]*Road) map[Player][]*Road {
	panic("unimplement")
}

// Set update board when board updated
func (imp *RoadBoard) Set(move Move) error {
	beginRow := move.Row
	beginCol := move.Col
	for i := 0; i < 5; i++ {
		row := beginRow - northToSouth.Row()*i
		col := beginCol - northToSouth.Col()*i
		road := imp.roadsPool[row][col]
		imp.removeRoad(imp.roads, road)
		road.Update(row, col, move.Player)
		imp.addRoad(imp.roads, road)
	}
	return nil
}

func (imp *RoadBoard) removeRoad(roadsLayer RoadsLayer, road *Road) {

}

func (imp *RoadBoard) addRoad(roadsLayer RoadsLayer, road *Road) {

}

// Regret move back n step
func (imp *RoadBoard) Regret(step int) error {
	panic("unimplement")
}

// GetPosInLiveThree get live three position in three road
func (imp *RoadBoard) GetPosInLiveThree(player Player) []Move {
	moves := make([]Move, 0)
	for _, r := range imp.liveThreeRoads[player] {
		poses := r.getEmptyPos()
		for _, p := range poses {
			moves = append(moves, Move{Row: p.Row, Col: p.Col, Player: player})
		}

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
		poses := r.getEmptyPos()
		for _, p := range poses {
			moves = append(moves, Move{Row: p.Row, Col: p.Col, Player: player})
		}

	}
	return moves
}

// Road road is continuous five position in specific direction
type Road struct {
}

func (imp *Road) getEmptyPos() []Pos {
	panic("unimplement")
}

// Update update road
func (imp *Road) Update(row, col int, player Player) {

}

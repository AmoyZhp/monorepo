package executor

import "fmt"

// RoadsBucket roads layer
type RoadsBucket map[Player][][]*Road

// RemoveRoad remove road. this will disrupt the order of the slice
func (imp RoadsBucket) RemoveRoad(road *Road) {
	roadsPointer := imp[road.BelongTo()][road.CountBelongPieces()]
	lastRoad := roadsPointer[len(roadsPointer)]
	lastRoad.SetIndex(road.Index())
	roadsPointer[road.Index()] = lastRoad
	imp[road.BelongTo()][road.CountBelongPieces()] = roadsPointer[:len(roadsPointer)-1]
}

// AddRoad add road to the end
func (imp RoadsBucket) AddRoad(road *Road) {
	temp := imp[road.BelongTo()][road.CountBelongPieces()]
	road.SetIndex(len(temp))
	imp[road.BelongTo()][road.CountBelongPieces()] = append(temp, road)
}

// RoadBoard board to store roads
type RoadBoard struct {
	roadsPool      [][]*Road
	roadsBucket    RoadsBucket
	liveThreeRoads map[Player][]*Road
}

// NewRoadBoard new road board
func NewRoadBoard() RoadBoard {

	roadsPool := newRoadsPool()
	roads := newRoads(roadsPool)
	liveThreeRoads := newLiveThreeRoads(roads)
	return RoadBoard{
		roadsPool:      roadsPool,
		roadsBucket:    roads,
		liveThreeRoads: liveThreeRoads,
	}
}

func newRoadsPool() [][]*Road {
	panic("unimplement")
}

func newRoads(roadsPool [][]*Road) RoadsBucket {
	panic("unimplement")
}

func newLiveThreeRoads(RoadsBucket) map[Player][]*Road {
	panic("unimplement")
}

// Set update board when board updated
func (imp *RoadBoard) Set(move Move) error {
	return imp.update(move)
}

// Regret move back step
func (imp *RoadBoard) Regret(move Move) error {
	return imp.update(move)
}

func (imp *RoadBoard) update(move Move) error {
	beginRow := move.Row
	beginCol := move.Col
	for i := 0; i < 5; i++ {
		row := beginRow - northToSouth.Row()*i
		col := beginCol - northToSouth.Col()*i
		road := imp.roadsPool[row][col]
		imp.roadsBucket.RemoveRoad(road)
		road.Update(row, col, move.Player)
		imp.roadsBucket.AddRoad(road)
	}
	return nil
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
	for _, r := range imp.roadsBucket[player][roadNum] {
		poses := r.getEmptyPos()
		for _, p := range poses {
			moves = append(moves, Move{Row: p.Row, Col: p.Col, Player: player})
		}

	}
	return moves
}

// Road road is continuous five position in specific direction
type Road struct {
	belong     Player
	posArr     []*Move
	unEmptyCnt int
	index      int
}

func (imp *Road) getEmptyPos() []Pos {
	pos := make([]Pos, 0)
	for _, p := range imp.posArr {
		if p.Player == EMPTY {
			pos = append(pos, Pos{Row: p.Row, Col: p.Col})
		}
	}
	return pos
}

// Update update road
func (imp *Road) Update(row, col int, player Player) {
	for _, p := range imp.posArr {
		if p.Row == row && p.Col == col {
			if p.Player != EMPTY {
				// TODO should occur error
				fmt.Println("road position has taken")
				return
			}
			p.Player = player
			return
		}
	}
	fmt.Println("pass wrong position to road update")
}

// BelongTo player road belong to
func (imp *Road) BelongTo() Player {
	return imp.belong
}

// CountPieces count player pieces
func (imp *Road) CountPieces(player Player) int {
	cnt := 0
	for _, p := range imp.posArr {
		if p.Player == player {
			cnt++
		}
	}
	return cnt
}

// CountBelongPieces count pieces of road belong to
func (imp *Road) CountBelongPieces() int {
	return imp.CountPieces(imp.BelongTo())
}

// CountAll count unempty position
func (imp *Road) CountAll() int {
	return imp.unEmptyCnt
}

// Index return index in roads bucket
func (imp *Road) Index() int {
	return imp.index
}

// SetIndex set index come from roads bucket
func (imp *Road) SetIndex(index int) {
	imp.index = index
}

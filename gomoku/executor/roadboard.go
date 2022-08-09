package executor

// NewRoadBoard new road board
func NewRoadBoard() RoadBoard {
	roadPool := newRoadPool()
	roadBucket := newRoadBucket(roadPool)
	return RoadBoard{
		roadsPool:   roadPool,
		roadsBucket: roadBucket,
	}
}

// RoadBoard board to store roads
type RoadBoard struct {
	roadsPool   [][][]*Road
	roadsBucket RoadBucket
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
	for _, dire := range directions {
		for i := 0; i < 5; i++ {
			row := beginRow - dire.Row()*i
			col := beginCol - dire.Col()*i
			road := imp.roadsPool[row][col][dire.Enum()]
			imp.roadsBucket.RemoveRoad(road)
			road.Update(row, col, move.Player)
			imp.roadsBucket.AddRoad(road)
		}
	}
	return nil
}

// GetPosInLiveThree get live three position in three road
func (imp *RoadBoard) GetPosInLiveThree(player Player) []Move {
	moves := make([]Move, 0)
	// TODO implement got pos in live three
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

// RoadBucket roads layer, player, road num, road
type RoadBucket map[Player][][]*Road

// RemoveRoad remove road. this will disrupt the order of the slice
func (imp RoadBucket) RemoveRoad(road *Road) {
	roadsPointer := imp[road.BelongTo()][road.CountBelongPieces()]
	lastRoad := roadsPointer[len(roadsPointer)]
	lastRoad.SetIndex(road.Index())
	roadsPointer[road.Index()] = lastRoad
	imp[road.BelongTo()][road.CountBelongPieces()] = roadsPointer[:len(roadsPointer)-1]
}

// AddRoad add road to the end
func (imp RoadBucket) AddRoad(road *Road) {
	temp := imp[road.BelongTo()][road.CountBelongPieces()]
	road.SetIndex(len(temp))
	imp[road.BelongTo()][road.CountBelongPieces()] = append(temp, road)
}

func newRoadPool() [][][]*Road {
	roadPool := make([][][]*Road, GomokuRow)
	for i := 0; i < GomokuRow; i++ {
		roadPool[i] = make([][]*Road, GomokuCol)
		for j := 0; j < GomokuCol; j++ {
			roadPool[i][j] = make([]*Road, 4)
			for _, dire := range directions {
				roadPool[i][j][dire.Enum()] = newRoad(i, j, dire)
			}
		}
	}
	return roadPool
}

func newRoadBucket(roadPool [][][]*Road) RoadBucket {
	var roadBucket RoadBucket
	roadBucket[EMPTY] = make([][]*Road, 5)
	roadBucket[BLACK] = make([][]*Road, 5)
	roadBucket[WHITE] = make([][]*Road, 5)
	roadBucket[NOMAN] = make([][]*Road, 5)
	for i := 0; i < 5; i++ {
		roadBucket[EMPTY][i] = make([]*Road, 0)
		roadBucket[BLACK][i] = make([]*Road, 0)
		roadBucket[WHITE][i] = make([]*Road, 0)
		roadBucket[NOMAN][i] = make([]*Road, 0)
	}
	for i := 0; i < GomokuRow; i++ {
		for j := 0; j < GomokuCol; j++ {
			for _, dire := range directions {
				road := roadPool[i][j][dire.Enum()]
				roadBucket[road.BelongTo()][road.CountBelongPieces()] = append(
					roadBucket[road.BelongTo()][road.CountBelongPieces()],
					road,
				)
			}
		}
	}
	return roadBucket
}

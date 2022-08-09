package executor

import "fmt"

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

// RoadBoard board to store roads
type RoadBoard struct {
	roadsPool   [][][]*Road
	roadsBucket RoadBucket
}

// NewRoadBoard new road board
func NewRoadBoard() RoadBoard {
	roadPool := newRoadPool()
	roadBucket := newRoadBucket(roadPool)
	return RoadBoard{
		roadsPool:   roadPool,
		roadsBucket: roadBucket,
	}
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

func newRoad(row, col int, dire direction) *Road {
	if row+dire.Row()*5 > GomokuRow || col*dire.Col()*5 > GomokuCol {
		return &Road{legal: false, belong: NOMAN}
	}
	pos := make([]*Move, 5)
	for i := 0; i < 5; i++ {
		pos = append(pos, &Move{Row: row * dire.Row() * i, Col: col * dire.Col() * i, Player: EMPTY})
	}
	playCnt := make(map[Player]int)
	playCnt[EMPTY] = 5
	road := &Road{
		legal:     true,
		belong:    EMPTY,
		posArr:    pos,
		playerCnt: playCnt,
	}
	return road
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

// Road road is continuous five position in specific direction
type Road struct {
	belong     Player
	posArr     []*Move
	unEmptyCnt int
	index      int
	playerCnt  map[Player]int
	legal      bool
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
			imp.playerCnt[p.Player]--
			p.Player = player
			imp.playerCnt[p.Player]++
			imp.updateBelong()
			return
		}
	}
	fmt.Println("pass wrong position to road update")
}

func (imp *Road) updateBelong() {
	maxCnt := -1
	belong := EMPTY
	for player, cnt := range imp.playerCnt {
		if cnt > maxCnt {
			belong = player
			maxCnt = cnt
		}
	}
	imp.belong = belong
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

// Legal it is a legal road
func (imp *Road) Legal() bool {
	return imp.legal
}

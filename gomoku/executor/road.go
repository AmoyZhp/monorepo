package executor

import "fmt"

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

package executor

// Player game player
type Player int

// player num
const (
	EMPTY Player = 0
	BLACK Player = 1
	WHITE Player = 2
	NOMAN Player = 100
)

func nextPlayer(player Player) Player {
	if player == WHITE {
		return BLACK
	}
	return WHITE
}

// Move chess act
type Move struct {
	Col    int    `json:"col"`
	Row    int    `json:"row"`
	Player Player `json:"player"`
}

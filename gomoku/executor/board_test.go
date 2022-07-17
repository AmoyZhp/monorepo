package executor

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGomokuBoard_Set(t *testing.T) {
	Convey("TestGomokuBoard_Set", t, func() {
		type args struct {
			move Move
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				"normal set",
				args{
					Move{7, 7, BLACK},
				},
				false,
			},
			{
				"invalid player",
				args{
					Move{7, 7, BLACK},
				},
				false,
			},
			{
				"invalid row",
				args{
					Move{19, 7, BLACK},
				},
				true,
			},
			{
				"invalid col",
				args{
					Move{7, 19, BLACK},
				},
				true,
			},
			{
				"positio is taken",
				args{
					Move{8, 8, BLACK},
				},
				true,
			},
		}
		for _, tt := range tests {
			Convey(tt.name, func() {
				imp := NewGomokuBoard()
				imp.Set(Move{8, 8, BLACK})
				err := imp.Set(tt.args.move)
				So(err != nil, ShouldEqual, tt.wantErr)
			})
		}
	})
}

func TestGomokuBoard_Reset(t *testing.T) {
	Convey("TestGomokuBoard_Reset", t, func() {
		Convey("reset", func() {
			innerBoard := make([][]Player, GomokuRow)
			for i := 0; i < GomokuRow; i++ {
				innerBoard[i] = make([]Player, GomokuCol)
			}
			historyMove := make([]Move, GomokuCol*GomokuRow)
			board := &GomokuBoard{
				board:       innerBoard,
				historyMove: historyMove,
				maxCol:      GomokuCol,
				maxRow:      GomokuRow,
				timeline:    0,
			}
			board.Set(Move{7, 7, BLACK})
			board.Reset()
			player, err := board.GetPlayerAtPos(7, 7)
			So(player, ShouldEqual, EMPTY)
			So(err, ShouldBeNil)
			So(board.timeline, ShouldEqual, 0)
			So(board.historyMove, ShouldResemble, historyMove)

		})

	})
}

func TestGomokuBoard_GetPlayerAtPos(t *testing.T) {
	Convey("TestGomokuBoard_GetPlayerAtPos", t, func() {
		type args struct {
			row int
			col int
		}
		tests := []struct {
			name    string
			args    args
			want    Player
			wantErr bool
		}{
			{
				"normal",
				args{7, 7},
				BLACK,
				false,
			},
			{
				"invalid row",
				args{
					19, 7,
				},
				EMPTY,
				true,
			},
			{
				"invalid col",
				args{
					7, 19,
				},
				EMPTY,
				true,
			},
		}
		for _, tt := range tests {
			Convey(tt.name, func() {
				imp := NewGomokuBoard()
				imp.Set(Move{7, 7, BLACK})
				got, err := imp.GetPlayerAtPos(tt.args.row, tt.args.col)
				So(got, ShouldEqual, tt.want)
				So(err != nil, ShouldEqual, tt.wantErr)
			})
		}
	})
}

func TestGomokuBoard_IsEnd(t *testing.T) {
	Convey("TestGomokuBoard_IsEnd", t, func() {
		Convey("not end", func() {
			board := NewGomokuBoard()
			So(board.IsEnd(), ShouldBeFalse)
		})
		Convey("end", func() {
			board := NewGomokuBoard()
			player := BLACK
			for i := 0; i < GomokuRow; i++ {
				for j := 0; j < GomokuCol; j++ {
					board.Set(Move{i, j, player})
					player = nextPlayer(player)
				}
			}
			So(board.IsEnd(), ShouldBeTrue)
		})
	})
}

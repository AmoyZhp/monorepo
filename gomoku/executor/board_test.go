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
				"invalid colum",
				args{
					Move{19, 7, BLACK},
				},
				true,
			},
			{
				"invalid row",
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

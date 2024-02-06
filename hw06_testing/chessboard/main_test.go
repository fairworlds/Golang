package chessboard

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc, want             string
		countColumn, firstStep int
	}{
		{
			desc:        "OneBlack",
			countColumn: 8,
			firstStep:   0,
			want:        "# # # # ",
		},
		{
			desc:        "OneWhite",
			countColumn: 9,
			firstStep:   1,
			want:        " # # # # ",
		},
		{
			desc:        "OneBlack1",
			countColumn: 5,
			firstStep:   0,
			want:        "# # #",
		},
		{
			desc:        "OneWhite2",
			countColumn: 6,
			firstStep:   3,
			want:        " # # #",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := CreateString(tC.countColumn, tC.firstStep)
			// assert.Equal(t, tC.want, got)
			if got != tC.want {
				t.Errorf("got %s, want %s", got, tC.want)
			}
		})
	}
}

func TestChessBoard(t *testing.T) {
	testCases := []struct {
		desc        string
		row, column int
		want        string
		er          error
	}{
		{
			desc:   "Goodblack",
			row:    5,
			column: 5,
			want:   "\n# # #\n # # \n# # #\n # # \n# # #",
			er:     nil,
		},
		{
			desc:   "Zero",
			row:    0,
			column: 5,
			want:   "",
			er:     errors.New("значение не может быть меньше или равно нулю"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := ChessBoard(tC.row, tC.column)
			if err != nil {
				assert.Equal(t, tC.er, err)
			}
			assert.Equal(t, tC.want, got)
		})
	}
}

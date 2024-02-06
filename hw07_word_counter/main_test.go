package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordSplit(t *testing.T) {
	testCases := []struct {
		desc string
		row  string
		er   error
		want []string
	}{
		{
			desc: "its good",
			row:  "Go go the server",
			er:   nil,
			want: []string{"Go", "go", "the", "server"},
		},
		{
			desc: "Test failed",
			row:  "",
			er:   errors.New("empty string"),
			want: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := WordSplit(tC.row)
			if err != nil {
				assert.Equal(t, tC.er, err)
			}
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestCountWorld(t *testing.T) {
	testCases := []struct {
		desc string
		row  []string
		want map[string]int
	}{
		{
			desc: "its good",
			row:  []string{"Go", "go", "the", "server"},
			want: map[string]int{
				"server": 1,
				"the":    1,
				"go":     2,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := countWorld(tC.row)
			assert.Equal(t, tC.want, got)
		})
	}
}

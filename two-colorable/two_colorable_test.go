package two_colorable

import (
	"strconv"
	"testing"
)

var stages = []struct {
	g      [][]int
	expect bool
}{
	{
		g: [][]int{
			{1, 2},
			{0, 2},
			{0, 1},
		},
	},
}

func TestTwoColorable(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := TwoColorable(stg.g)
			if r != stg.expect {
				println(r)
				t.FailNow()
			}
		})
	}
}

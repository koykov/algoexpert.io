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
		g:      [][]int{{1}, {0}},
		expect: true,
	},
	{g: [][]int{{0}}},
	{
		g: [][]int{
			{1},
			{0, 2},
			{1},
		},
		expect: true,
	},
	{
		g: [][]int{
			{1, 2, 3},
			{0},
			{0},
			{0},
		},
		expect: true,
	},
	{
		g: [][]int{
			{2},
			{2, 3},
			{0, 1},
			{1},
		},
		expect: true,
	},
	{
		g: [][]int{
			{1, 4},
			{0, 2, 3},
			{1, 4},
			{1},
			{0, 2},
		},
		expect: true,
	},
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
				t.FailNow()
			}
		})
	}
}

func BenchmarkTwoColorable(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := stages[i%len(stages)]
		TwoColorable(stg.g)
	}
}

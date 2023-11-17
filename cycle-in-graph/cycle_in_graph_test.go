package cycle_in_graph

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
			{1, 3},
			{2, 3, 4},
			{0},
			{},
			{2, 5},
			{},
		},
		expect: true,
	},
	{
		g: [][]int{
			{1, 2},
			{2},
			{},
		},
	},
	{
		g: [][]int{
			{},
			{0, 2},
			{0, 3},
			{0, 4},
			{0, 5},
			{0},
		},
	},
}

func TestCycleInGraph(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := CycleInGraph(stg.g)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkCycleInGraph(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		CycleInGraph(stg.g)
	}
}

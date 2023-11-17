package minimum_passes_of_matrix

import (
	"strconv"
	"testing"
)

var stages = []struct {
	m      [][]int
	expect int
}{
	{
		m: [][]int{
			{0, -1, -3, 2, 0},
			{1, -2, -5, -1, -3},
			{3, 0, 0, -4, -1},
		},
		expect: 3,
	},
	{
		m: [][]int{
			{1, 0, 0, -2, -3},
			{-4, -5, -6, -2, -1},
			{0, 0, 0, 0, -1},
			{-1, 0, 3, 0, 3},
		},
		expect: -1,
	},
}

func TestRemoveIslands(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := MinimumPassesOfMatrix(stg.m)
			if r != stg.expect {
				println(r)
				t.FailNow()
			}
		})
	}
}

func BenchmarkRemoveIslands(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := stages[i%len(stages)]
		MinimumPassesOfMatrix(stg.m)
	}
}

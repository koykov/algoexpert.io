package remove_islands

import (
	"fmt"
	"strconv"
	"testing"
)

var stages = []struct {
	m, e [][]int
}{
	{
		m: [][]int{
			{1, 0, 0, 0, 0, 0},
			{0, 1, 0, 1, 1, 1},
			{0, 0, 1, 0, 1, 0},
			{1, 1, 0, 0, 1, 0},
			{1, 0, 1, 1, 0, 0},
			{1, 0, 0, 0, 0, 1},
		},
		e: [][]int{
			{1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 1},
			{0, 0, 0, 0, 1, 0},
			{1, 1, 0, 0, 1, 0},
			{1, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 1},
		},
	},
	{
		m: [][]int{
			{1, 0, 0, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 0, 1, 1, 0},
		},
		e: [][]int{
			{1, 0, 0, 1, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 1, 1, 0},
		},
	},
	{
		m: [][]int{
			{1, 1, 1, 1, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 1, 1, 1},
			{1, 0, 1, 0, 1},
		},
		e: [][]int{
			{1, 1, 1, 1, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 1, 1, 1},
			{1, 0, 1, 0, 1},
		},
	},
}

func TestRemoveIslands(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := RemoveIslands(stg.m)
			rm, em := fmt.Sprintf("%v", r), fmt.Sprintf("%v", stg.e)
			if rm != em {
				println(mprint(stg.m))
				t.FailNow()
			}
		})
	}
}

func BenchmarkRemoveIslands(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := stages[i%len(stages)]
		RemoveIslands(stg.m)
	}
}

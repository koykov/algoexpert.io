package river_sizes

import (
	"fmt"
	"testing"
)

var stages = []struct {
	m [][]int
	r []int
}{
	{
		m: [][]int{
			{0, 0, 0, 0, 0},
			{0, 1, 0, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 1, 1, 1, 0},
			{0, 0, 0, 0, 0},
		},
		r: []int{7},
	},
	{
		m: [][]int{
			{1, 0, 0, 1, 0},
			{1, 0, 1, 0, 0},
			{0, 0, 1, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 1, 1, 0},
		},
		r: []int{2, 1, 5, 2, 2},
	},
	{
		m: [][]int{{0}},
		r: []int{},
	},
	{
		m: [][]int{{1}},
		r: []int{1},
	},
	{
		m: [][]int{
			{1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0},
		},
		r: []int{3, 2, 1},
	},
	{
		m: [][]int{
			{1, 0, 0, 1},
			{1, 0, 1, 0},
			{0, 0, 1, 0},
			{1, 0, 1, 0},
		},
		r: []int{2, 1, 3, 1},
	},
	{
		m: [][]int{
			{1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0},
			{1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0},
			{0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 1},
			{1, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0},
			{1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0, 1},
		},
		r: []int{2, 1, 21, 5, 2, 1},
	},
	{
		m: [][]int{
			{1, 0, 0, 0, 0, 0, 1},
			{0, 1, 0, 0, 0, 1, 0},
			{0, 0, 1, 0, 1, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 0, 1, 0},
			{1, 0, 0, 0, 0, 0, 1},
		},
		r: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	},
	{
		m: [][]int{
			{1, 0, 0, 0, 0, 0, 1},
			{0, 1, 0, 0, 0, 1, 0},
			{0, 0, 1, 0, 1, 0, 0},
			{0, 0, 1, 1, 1, 0, 0},
			{0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 0, 1, 0},
			{1, 0, 0, 0, 0, 0, 1},
		},
		r: []int{1, 1, 1, 1, 7, 1, 1, 1, 1},
	},
	{
		m: [][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
		r: []int{},
	},
	{
		m: [][]int{
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
		},
		r: []int{49},
	},
	{
		m: [][]int{
			{1, 1, 0, 0, 0, 0, 1, 1},
			{1, 0, 1, 1, 1, 1, 0, 1},
			{0, 1, 1, 0, 0, 0, 1, 1},
		},
		r: []int{3, 5, 6},
	},
	{
		m: [][]int{
			{1, 1, 0},
			{1, 0, 1},
			{1, 1, 1},
			{1, 1, 0},
			{1, 0, 1},
			{0, 1, 0},
			{1, 0, 0},
			{1, 0, 0},
			{0, 0, 0},
			{1, 0, 0},
			{1, 0, 1},
			{1, 1, 1},
		},
		r: []int{10, 1, 1, 2, 6},
	},
}

func TestRiverSizes(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run(fmt.Sprintf("dfs/%d", i), func(t *testing.T) {
			c := RiverSizes(stage.m)
			if !assert(c, stage.r) {
				t.FailNow()
			}
		})
		t.Run(fmt.Sprintf("dfs optimized/%d", i), func(t *testing.T) {
			c := RiverSizesOptimized(stage.m)
			if !assert(c, stage.r) {
				t.FailNow()
			}
		})
	}
}

func assert(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkRiverSizes(b *testing.B) {
	b.Run("dfs", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			stage := &stages[i%len(stages)]
			RiverSizes(stage.m)
		}
	})
	b.Run("dfs optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			stage := &stages[i%len(stages)]
			RiverSizesOptimized(stage.m)
		}
	})
}

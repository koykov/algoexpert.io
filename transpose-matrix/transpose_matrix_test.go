package transpose_matrix

import (
	"fmt"
	"strconv"
	"testing"
)

var stages = []struct {
	matrix, expect [][]int
}{
	{
		matrix: [][]int{
			{1, 2},
			{3, 4},
			{5, 6},
		},
		expect: [][]int{
			{1, 3, 5},
			{2, 4, 6},
		},
	},
	{
		matrix: [][]int{{1}},
		expect: [][]int{{1}},
	},
	{
		matrix: [][]int{{1}, {-1}},
		expect: [][]int{{1, -1}},
	},
	{
		matrix: [][]int{{1, 2, 3}},
		expect: [][]int{{1}, {2}, {3}},
	},
	{
		matrix: [][]int{{1}, {2}, {3}},
		expect: [][]int{{1, 2, 3}},
	},
	{
		matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
		},
		expect: [][]int{
			{1, 4},
			{2, 5},
			{3, 6},
		},
	},
	{
		matrix: [][]int{
			{0, 0, 0},
			{1, 1, 1},
		},
		expect: [][]int{
			{0, 1},
			{0, 1},
			{0, 1},
		},
	},
	{
		matrix: [][]int{
			{0, 1},
			{0, 1},
			{0, 1},
		},
		expect: [][]int{
			{0, 0, 0},
			{1, 1, 1},
		},
	},
	{
		matrix: [][]int{
			{0, 0, 0},
			{0, 0, 0},
		},
		expect: [][]int{
			{0, 0},
			{0, 0},
			{0, 0},
		},
	},
	{
		matrix: [][]int{
			{1, 4},
			{2, 5},
			{3, 6},
		},
		expect: [][]int{
			{1, 2, 3},
			{4, 5, 6},
		},
	},
	{
		matrix: [][]int{
			{-7, -7},
			{100, 12},
			{-33, 17},
		},
		expect: [][]int{
			{-7, 100, -33},
			{-7, 12, 17},
		},
	},
	{
		matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		expect: [][]int{
			{1, 4, 7},
			{2, 5, 8},
			{3, 6, 9},
		},
	},
	{
		matrix: [][]int{
			{1, 4, 7},
			{2, 5, 8},
			{3, 6, 9},
		},
		expect: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	},
	{
		matrix: [][]int{
			{5, 6, 3, -3, 12},
			{-3, 6, 5, 2, -1},
			{0, 0, 3, 12, 3},
		},
		expect: [][]int{
			{5, -3, 0},
			{6, 6, 0},
			{3, 5, 3},
			{-3, 2, 12},
			{12, -1, 3},
		},
	},
	{
		matrix: [][]int{
			{0, -1, -2, -3},
			{4, 5, 6, 7},
			{2, 3, -2, -3},
			{42, 100, 30, -42},
		},
		expect: [][]int{
			{0, 4, 2, 42},
			{-1, 5, 3, 100},
			{-2, 6, -2, 30},
			{-3, 7, -3, -42},
		},
	},
	{
		matrix: [][]int{
			{1234, 6935, 4205},
			{-23459, 314159, 0},
			{100, 3, 987654},
		},
		expect: [][]int{
			{1234, -23459, 100},
			{6935, 314159, 3},
			{4205, 0, 987654},
		},
	},
}

func TestTransposeMatrix(t *testing.T) {
	eq := func(a, b [][]int) bool {
		return fmt.Sprintf("%#v", a) == fmt.Sprintf("%#v", b)
	}
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := TransposeMatrix(stage.matrix)
			if !eq(r, stage.expect) {
				t.FailNow()
			}
		})
	}
}

func BenchmarkTransposeMatrix(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stage := &stages[i%len(stages)]
		TransposeMatrix(stage.matrix)
	}
}

package spiral_traverse

import (
	"fmt"
	"strconv"
	"testing"
)

var stages = []struct {
	matrix [][]int
	expect []int
}{
	{
		matrix: [][]int{
			{1, 2, 3, 4},
			{12, 13, 14, 5},
			{11, 16, 15, 6},
			{10, 9, 8, 7},
		},
		expect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
	},
	{
		matrix: [][]int{{1}},
		expect: []int{1},
	},
	{
		matrix: [][]int{
			{1, 2},
			{4, 3},
		},
		expect: []int{1, 2, 3, 4},
	},
	{
		matrix: [][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		},
		expect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		matrix: [][]int{
			{19, 32, 33, 34, 25, 8},
			{16, 15, 14, 13, 12, 11},
			{18, 31, 36, 35, 26, 9},
			{1, 2, 3, 4, 5, 6},
			{20, 21, 22, 23, 24, 7},
			{17, 30, 29, 28, 27, 10},
		},
		expect: []int{19, 32, 33, 34, 25, 8, 11, 9, 6, 7, 10, 27, 28, 29, 30, 17, 20, 1, 18, 16, 15, 14, 13, 12, 26, 5, 24, 23, 22, 21, 2, 31, 36, 35, 4, 3},
	},
	{
		matrix: [][]int{
			{4, 2, 3, 6, 7, 8, 1, 9, 5, 10},
			{12, 19, 15, 16, 20, 18, 13, 17, 11, 14},
		},
		expect: []int{4, 2, 3, 6, 7, 8, 1, 9, 5, 10, 14, 11, 17, 13, 18, 20, 16, 15, 19, 12},
	},
	{
		matrix: [][]int{
			{27, 12, 35, 26},
			{25, 21, 94, 11},
			{19, 96, 43, 56},
			{55, 36, 10, 18},
			{96, 83, 31, 94},
			{93, 11, 90, 16},
		},
		expect: []int{27, 12, 35, 26, 11, 56, 18, 94, 16, 90, 11, 93, 96, 55, 19, 25, 21, 94, 43, 10, 31, 83, 36, 96},
	},
	{
		matrix: [][]int{
			{1, 2, 3, 4},
			{10, 11, 12, 5},
			{9, 8, 7, 6},
		},
		expect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
	},
	{
		matrix: [][]int{
			{1, 2, 3},
			{12, 13, 4},
			{11, 14, 5},
			{10, 15, 6},
			{9, 8, 7},
		},
		expect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
	},
	{
		matrix: [][]int{
			{1, 11},
			{2, 12},
			{3, 13},
			{4, 14},
			{5, 15},
			{6, 16},
			{7, 17},
			{8, 18},
			{9, 19},
			{10, 20},
		},
		expect: []int{1, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 10, 9, 8, 7, 6, 5, 4, 3, 2},
	},
	{
		matrix: [][]int{{1, 3, 2, 5, 4, 7, 6}},
		expect: []int{1, 3, 2, 5, 4, 7, 6},
	},
	{
		matrix: [][]int{
			{1},
			{3},
			{2},
			{5},
			{4},
			{7},
			{6},
		},
		expect: []int{1, 3, 2, 5, 4, 7, 6},
	},
}

func TestSpiralTraverse(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			stg := &stages[i]
			r := SpiralTraverse(stg.matrix)
			if fmt.Sprintf("%#v", r) != fmt.Sprintf("%#v", stg.expect) {
				t.Log(fmt.Sprintf("%#v", r))
				t.Log(fmt.Sprintf("%#v", stg.expect))
				t.FailNow()
			}
		})
	}
}

func BenchmarkSpiralTraverse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		SpiralTraverse(stg.matrix)
	}
}

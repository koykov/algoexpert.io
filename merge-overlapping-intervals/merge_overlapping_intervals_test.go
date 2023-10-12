package merge_overlapping_intervals

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
			{3, 5},
			{4, 7},
			{6, 8},
			{9, 10},
		},
		expect: [][]int{{1, 2}, {3, 8}, {9, 10}},
	},
	{
		matrix: [][]int{
			{100, 105},
			{1, 104},
		},
		expect: [][]int{{1, 105}},
	},
	{
		matrix: [][]int{
			{89, 90},
			{-10, 20},
			{-50, 0},
			{70, 90},
			{90, 91},
			{90, 95},
		},
		expect: [][]int{{-50, 20}, {70, 95}},
	},
}

func TestMergeOverlappingIntervals(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			stg := &stages[i]
			r := MergeOverlappingIntervals(stg.matrix)
			if fmt.Sprintf("%#v", r) != fmt.Sprintf("%#v", stg.expect) {
				t.FailNow()
			}
		})
	}
}

func BenchmarkMergeOverlappingIntervals(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		MergeOverlappingIntervals(stg.matrix)
	}
}

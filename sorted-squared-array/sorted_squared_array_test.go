package sorted_squared_array

import (
	"strconv"
	"testing"
)

var stages = []struct {
	array, expect []int
}{
	{
		array:  []int{1, 2, 3, 5, 6, 8, 9},
		expect: []int{1, 4, 9, 25, 36, 64, 81},
	},
	{
		array:  []int{1},
		expect: []int{1},
	},
	{
		array:  []int{1, 2},
		expect: []int{1, 4},
	},
	{
		array:  []int{1, 2, 3, 4, 5},
		expect: []int{1, 4, 9, 16, 25},
	},
	{
		array:  []int{10},
		expect: []int{100},
	},
	{
		array:  []int{-1},
		expect: []int{1},
	},
	{
		array:  []int{-2, -1},
		expect: []int{1, 4},
	},
	{
		array:  []int{-5, -4, -3, -2, -1},
		expect: []int{1, 4, 9, 16, 25},
	},
	{
		array:  []int{-10},
		expect: []int{100},
	},
	{
		array:  []int{-10, -5, 0, 5, 10},
		expect: []int{0, 25, 25, 100, 100},
	},
	{
		array:  []int{-7, -3, 1, 9, 22, 30},
		expect: []int{1, 9, 49, 81, 484, 900},
	},
	{
		array:  []int{-50, -13, -2, -1, 0, 0, 1, 1, 2, 3, 19, 20},
		expect: []int{0, 0, 1, 1, 1, 4, 4, 9, 169, 361, 400, 2500},
	},
}

func TestTwoNumberSum(t *testing.T) {
	eq := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		if len(a) == 0 && len(b) == 0 {
			return true
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := SortedSquaredArray(stage.array)
			if !eq(r, stage.expect) {
				t.FailNow()
			}
		})
	}
}

func BenchmarkSortedSquaredArray(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stage := &stages[i%len(stages)]
		SortedSquaredArray(stage.array)
	}
}

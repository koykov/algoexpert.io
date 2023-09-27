package two_number_sum

import (
	"strconv"
	"testing"
)

var stages = []struct {
	array  []int
	target int
	expect []int
}{
	// {
	// 	array:  []int{3, 5, -4, 8, 11, 1, -1, 6},
	// 	target: 10,
	// 	expect: []int{11, -1},
	// },
	// {
	// 	array:  []int{4, 6},
	// 	target: 10,
	// 	expect: []int{4, 6},
	// },
	// {
	// 	array:  []int{4, 6, 1},
	// 	target: 5,
	// 	expect: []int{4, 1},
	// },
	// {
	// 	array:  []int{4, 6, 1, -3},
	// 	target: 3,
	// 	expect: []int{6, -3},
	// },
	// {
	// 	array:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	target: 17,
	// 	expect: []int{8, 9},
	// },
	// {
	// 	array:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15},
	// 	target: 18,
	// 	expect: []int{3, 15},
	// },
	// {
	// 	array:  []int{-7, -5, -3, -1, 0, 1, 3, 5, 7},
	// 	target: -5,
	// 	expect: []int{-5, 0},
	// },
	// {
	// 	array:  []int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47},
	// 	target: 163,
	// 	expect: []int{210, -47},
	// },
	{
		array:  []int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47},
		target: 164,
		expect: []int{},
	},
	{
		array:  []int{14},
		target: 15,
		expect: []int{},
	},
	{
		array:  []int{15},
		target: 15,
		expect: []int{},
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
		return a[0] == b[0] && a[1] == b[1]
	}
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := TwoNumberSum(stage.array, stage.target)
			if !eq(r, stage.expect) {
				t.FailNow()
			}
		})
	}
}

func BenchmarkMaxProfitWithKTransactions(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stage := &stages[i%len(stages)]
		TwoNumberSum(stage.array, stage.target)
	}
}

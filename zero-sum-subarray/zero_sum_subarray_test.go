package zero_sum_subarray

import (
	"strconv"
	"testing"
)

var stages = []struct {
	arr    []int
	expect bool
}{
	{
		arr:    []int{},
		expect: false,
	},
	{
		arr:    []int{0},
		expect: true,
	},
	{
		arr:    []int{1},
		expect: false,
	},
	{
		arr:    []int{1, 2, 3},
		expect: false,
	},
	{
		arr:    []int{1, 1, 1},
		expect: false,
	},
	{
		arr:    []int{-1, -1, -1},
		expect: false,
	},
	{
		arr:    []int{0, 0, 0},
		expect: true,
	},
	{
		arr:    []int{1, 2, -2, 3},
		expect: true,
	},
	{
		arr:    []int{2, -2},
		expect: true,
	},
	{
		arr:    []int{-1, 2, 3, 4, -5, -3, 1, 2},
		expect: true,
	},
	{
		arr:    []int{-2, -3, -1, 2, 3, 4, -5, -3, 1, 2},
		expect: true,
	},
	{
		arr:    []int{2, 3, 4, -5, -3, 4, 5},
		expect: true,
	},
	{
		arr:    []int{2, 3, 4, -5, -3, 5, 5},
		expect: false,
	},
	{
		arr:    []int{1, 2, 3, 4, 0, 5, 6, 7},
		expect: true,
	},
	{
		arr:    []int{1, 2, 3, -2, -1},
		expect: true,
	},
	{
		arr:    []int{-8, -22, 104, 73, -120, 53, 22, -12, 1, 14, -90, 13, -22},
		expect: false,
	},
	{
		arr:    []int{-8, -22, 104, 73, -120, 53, 22, 20, 25, -12, 1, 14, -90, 13, -22},
		expect: true,
	},
}

func TestZeroSumSubarray(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			stg := &stages[i]
			r := ZeroSumSubarray(stg.arr)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkZeroSumSubarray(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		ZeroSumSubarray(stg.arr)
	}
}

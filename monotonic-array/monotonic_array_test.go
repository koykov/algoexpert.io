package monotonic_array

import (
	"strconv"
	"testing"
)

var stages = []struct {
	arr    []int
	expect bool
}{
	{
		arr:    []int{-1, -5, -10, -1100, -1101, -1102, -9001},
		expect: true,
	},
	{
		arr:    []int{-1, -5, -10, -1100, -900, -1101, -1102, -9001},
		expect: false,
	},
	{
		arr:    []int{1, 2, 0},
		expect: false,
	},
	{
		arr:    []int{1, 1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 7, 9, 10, 11},
		expect: false,
	},
	{
		arr:    []int{1, 1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 8, 9, 10, 11},
		expect: true,
	},
	{
		arr:    []int{-1, -1, -2, -3, -4, -5, -5, -5, -6, -7, -8, -7, -9, -10, -11},
		expect: false,
	},
	{
		arr:    []int{-1, -1, -2, -3, -4, -5, -5, -5, -6, -7, -8, -8, -9, -10, -11},
		expect: true,
	},
	{
		arr:    []int{-1, -1, -1, -1, -1, -1, -1, -1},
		expect: true,
	},
	{
		arr:    []int{1, 2, -1, -2, -5},
		expect: false,
	},
	{
		arr:    []int{-1, -5, 10},
		expect: false,
	},
	{
		arr:    []int{2, 2, 2, 1, 4, 5},
		expect: false,
	},
	{
		arr:    []int{1, 1, 1, 2, 3, 4, 1},
		expect: false,
	},
	{
		arr:    []int{1, 2, 3, 3, 2, 1},
		expect: false,
	},
}

func TestIsMonotonic(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			stg := &stages[i]
			r := IsMonotonic(stg.arr)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkIsMonotonic(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		IsMonotonic(stg.arr)
	}
}

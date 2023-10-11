package longest_peak

import (
	"fmt"
	"strconv"
	"testing"
)

var stages = []struct {
	arr    []int
	expect int
}{
	// {
	// 	arr:    []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3},
	// 	expect: 6,
	// },
	// {
	// 	arr:    []int{},
	// 	expect: 0,
	// },
	// {
	// 	arr:    []int{1, 3, 2},
	// 	expect: 3,
	// },
	// {
	// 	arr:    []int{1, 2, 3, 4, 5, 1},
	// 	expect: 6,
	// },
	// {
	// 	arr:    []int{5, 4, 3, 2, 1, 2, 1},
	// 	expect: 3,
	// },
	// {
	// 	arr:    []int{5, 4, 3, 2, 1, 2, 10, 12, -3, 5, 6, 7, 10},
	// 	expect: 5,
	// },
	// {
	// 	arr:    []int{5, 4, 3, 2, 1, 2, 10, 12},
	// 	expect: 0,
	// },
	// {
	// 	arr:    []int{1, 2, 3, 4, 5, 6, 10, 100, 1000},
	// 	expect: 0,
	// },
	// {
	// 	arr:    []int{1, 2, 3, 3, 2, 1},
	// 	expect: 0,
	// },
	// {
	// 	arr:    []int{1, 1, 3, 2, 1},
	// 	expect: 4,
	// },
	// {
	// 	arr:    []int{1, 2, 3, 2, 1, 1},
	// 	expect: 5,
	// },
	{
		arr:    []int{1, 1, 1, 2, 3, 10, 12, -3, -3, 2, 3, 45, 800, 99, 98, 0, -1, -1, 2, 3, 4, 5, 0, -1, -1},
		expect: 9,
	},
	{
		arr:    []int{1, 2, 3, 3, 4, 0, 10},
		expect: 3,
	},
}

func TestLongestPeak(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			stg := &stages[i]
			r := LongestPeak(stg.arr)
			if r != stg.expect {
				println(r, stg.expect, fmt.Sprintf("%#v", stg.arr))
				t.FailNow()
			}
		})
	}
}

func BenchmarkLongestPeak(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		LongestPeak(stg.arr)
	}
}

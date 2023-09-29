package find_three_largest_numbers

import (
	"strconv"
	"testing"
)

var stages = []struct {
	array  []int
	expect []int
}{
	{
		array:  []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7},
		expect: []int{18, 141, 541},
	},
	{
		array:  []int{55, 7, 8},
		expect: []int{7, 8, 55},
	},
	{
		array:  []int{55, 43, 11, 3, -3, 10},
		expect: []int{11, 43, 55},
	},
	{
		array:  []int{7, 8, 3, 11, 43, 55},
		expect: []int{11, 43, 55},
	},
	{
		array:  []int{55, 7, 8, 3, 43, 11},
		expect: []int{11, 43, 55},
	},
	{
		array:  []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		expect: []int{7, 7, 7},
	},
	{
		array:  []int{-1, -2, -3, -7, -17, -27, -18, -541, -8, -7, 7},
		expect: []int{-2, -1, 7},
	},
}

func TestFindThreeLargestNumbers(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := FindThreeLargestNumbers(stg.array)
			if r[0] != stg.expect[0] || r[1] != stg.expect[1] || r[2] != stg.expect[2] {
				t.FailNow()
			}
		})
	}
}

func BenchmarkFindThreeLargestNumbers(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		FindThreeLargestNumbers(stg.array)
	}
}

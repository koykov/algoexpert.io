package smallest_difference

import (
	"strconv"
	"testing"
)

var stages = []struct {
	a1, a2 []int
	expect []int
}{
	{
		a1:     []int{-1, 5, 10, 20, 28, 3},
		a2:     []int{26, 134, 135, 15, 17},
		expect: []int{28, 26},
	},
	{
		a1:     []int{-1, 5, 10, 20, 3},
		a2:     []int{26, 134, 135, 15, 17},
		expect: []int{20, 17},
	},
	{
		a1:     []int{10, 0, 20, 25},
		a2:     []int{1005, 1006, 1014, 1032, 1031},
		expect: []int{25, 1005},
	},
	{
		a1:     []int{10, 0, 20, 25, 2200},
		a2:     []int{1005, 1006, 1014, 1032, 1031},
		expect: []int{25, 1005},
	},
	{
		a1:     []int{10, 0, 20, 25, 2000},
		a2:     []int{1005, 1006, 1014, 1032, 1031},
		expect: []int{2000, 1032},
	},
	{
		a1:     []int{240, 124, 86, 111, 2, 84, 954, 27, 89},
		a2:     []int{1, 3, 954, 19, 8},
		expect: []int{954, 954},
	},
	{
		a1:     []int{0, 20},
		a2:     []int{21, -2},
		expect: []int{20, 21},
	},
	{
		a1:     []int{10, 1000},
		a2:     []int{-1441, -124, -25, 1014, 1500, 660, 410, 245, 530},
		expect: []int{1000, 1014},
	},
	{
		a1:     []int{10, 1000, 9124, 2142, 59, 24, 596, 591, 124, -123},
		a2:     []int{-1441, -124, -25, 1014, 1500, 660, 410, 245, 530},
		expect: []int{-123, -124},
	},
	{
		a1:     []int{10, 1000, 9124, 2142, 59, 24, 596, 591, 124, -123, 530},
		a2:     []int{-1441, -124, -25, 1014, 1500, 660, 410, 245, 530},
		expect: []int{530, 530},
	},
}

func TestSmallestDifference(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			stg := &stages[i]
			r := SmallestDifference(stg.a1, stg.a2)
			if r[0] != stg.expect[0] || r[1] != stg.expect[1] {
				t.FailNow()
			}
		})
	}
}

func BenchmarkSmallestDifference(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		SmallestDifference(stg.a1, stg.a2)
	}
}

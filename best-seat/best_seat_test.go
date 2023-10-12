package best_seat

import (
	"strconv"
	"testing"
)

var stages = []struct {
	arr    []int
	expect int
}{
	{
		arr:    []int{1},
		expect: -1,
	},
	{
		arr:    []int{1, 0, 1, 0, 0, 0, 1},
		expect: 4,
	},
	{
		arr:    []int{1, 0, 1},
		expect: 1,
	},
	{
		arr:    []int{1, 0, 0, 1},
		expect: 1,
	},
	{
		arr:    []int{1, 1, 1},
		expect: -1,
	},
	{
		arr:    []int{1, 0, 0, 1, 0, 0, 1},
		expect: 1,
	},
	{
		arr:    []int{1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		expect: 3,
	},
	{
		arr:    []int{1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		expect: 4,
	},
	{
		arr:    []int{1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1},
		expect: 4,
	},
	{
		arr:    []int{1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		expect: 13,
	},
	{
		arr:    []int{1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1},
		expect: 13,
	},
	{
		arr:    []int{1, 0, 0, 0, 1, 0, 0, 0, 0, 1},
		expect: 6,
	},
	{
		arr:    []int{1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1},
		expect: 3,
	},
	{
		arr:    []int{1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1},
		expect: 5,
	},
}

func TestBestSeat(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			stg := &stages[i]
			r := BestSeat(stg.arr)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkBestSeat(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		BestSeat(stg.arr)
	}
}

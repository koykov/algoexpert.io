package minimum_waiting_time

import (
	"strconv"
	"testing"
)

var stages = []struct {
	queries []int
	expect  int
}{
	{
		queries: []int{3, 2, 1, 2, 6},
		expect:  17,
	},
	{
		queries: []int{2, 1, 1, 1},
		expect:  6,
	},
	{
		queries: []int{1, 2, 4, 5, 2, 1},
		expect:  23,
	},
	{
		queries: []int{25, 30, 2, 1},
		expect:  32,
	},
	{
		queries: []int{1, 1, 1, 1, 1},
		expect:  10,
	},
	{
		queries: []int{7, 9, 2, 3},
		expect:  19,
	},
	{
		queries: []int{4, 3, 1, 1, 3, 2, 1, 8},
		expect:  45,
	},
	{
		queries: []int{2},
		expect:  0,
	},
	{
		queries: []int{7},
		expect:  0,
	},
	{
		queries: []int{8, 9},
		expect:  8,
	},
	{
		queries: []int{1, 9},
		expect:  1,
	}, {
		queries: []int{5, 4, 3, 2, 1},
		expect:  20,
	},
	{
		queries: []int{1, 1, 1, 4, 5, 6, 8, 1, 1, 2, 1},
		expect:  81,
	},
	{
		queries: []int{17, 4, 3},
		expect:  10,
	},
}

func TestMinimumWaitingTime(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := MinimumWaitingTime(stg.queries)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkMinimumWaitingTime(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		MinimumWaitingTime(stg.queries)
	}
}

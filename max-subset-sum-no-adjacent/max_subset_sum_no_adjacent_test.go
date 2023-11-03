package max_subset_sum_no_adjacent

import (
	"strconv"
	"testing"
)

var stages = []struct {
	arr    []int
	expect int
}{
	{
		arr:    []int{75, 105, 120, 75, 90, 135},
		expect: 330,
	},
	{
		arr:    []int{30, 25, 50, 55, 100, 120},
		expect: 205,
	},
}

func TestMaxSubsetSumNoAdjacent(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := MaxSubsetSumNoAdjacent(stg.arr)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkMaxSubsetSumNoAdjacent(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		MaxSubsetSumNoAdjacent(stg.arr)
	}
}

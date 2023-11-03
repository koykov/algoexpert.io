package min_number_of_coins_for_change

import (
	"strconv"
	"testing"
)

var stages = []struct {
	arr    []int
	n      int
	expect int
}{
	// {
	// 	arr:    []int{1, 5, 10},
	// 	n:      7,
	// 	expect: 3,
	// },
	{
		arr:    []int{2, 1},
		n:      3,
		expect: 2,
	},
}

func TestMinNumberOfCoinsForChange(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := MinNumberOfCoinsForChange(stg.n, stg.arr)
			if r != stg.expect {
				println(r)
				t.FailNow()
			}
		})
	}
}

func BenchmarkMinNumberOfCoinsForChange(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		MinNumberOfCoinsForChange(stg.n, stg.arr)
	}
}

package single_cycle_check

import (
	"strconv"
	"testing"
)

var stages = []struct {
	a      []int
	expect bool
}{
	{
		a:      []int{2, 3, 1, -4, -4, 2},
		expect: true,
	},
	{
		a:      []int{2, 2, -1},
		expect: true,
	},
	{
		a:      []int{1, 1, 1, 1, 2},
		expect: false,
	},
}

func TestSingleCycleCheck(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := HasSingleCycle(stg.a)
			if r != stg.expect {
				println(r)
				t.FailNow()
			}
		})
	}
}

func BenchmarkSingleCycle(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		HasSingleCycle(stg.a)
	}
}

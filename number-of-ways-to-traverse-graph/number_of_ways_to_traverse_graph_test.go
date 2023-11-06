package number_of_ways_to_traverse_graph

import (
	"strconv"
	"testing"
)

var stages = []struct {
	w, h, expect int
}{
	{
		w:      4,
		h:      3,
		expect: 10,
	},
}

func TestNumberOfWaysToTraverseGraph(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := NumberOfWaysToTraverseGraph(stg.w, stg.h)
			if r != stg.expect {
				println(r)
				t.FailNow()
			}
		})
	}
}

func BenchmarkNumberOfWaysToTraverseGraph(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		NumberOfWaysToTraverseGraph(stg.w, stg.h)
	}
}

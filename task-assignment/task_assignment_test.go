package task_assignment

import (
	"fmt"
	"strconv"
	"testing"
)

var stages = []struct {
	k      int
	tasks  []int
	expect [][]int
}{
	{
		k:     3,
		tasks: []int{1, 3, 5, 3, 1, 4},
		expect: [][]int{
			{0, 2},
			{4, 5},
			{1, 3},
		},
	},
}

func TestTaskAssignment(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := TaskAssignment(stg.k, stg.tasks)
			rf := fmt.Sprintf("%v", r)
			ef := fmt.Sprintf("%v", stg.expect)
			if rf != ef {
				println(rf)
				t.FailNow()
			}
		})
	}
}

func BenchmarkTaskAssignment(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := stages[i%len(stages)]
		TaskAssignment(stg.k, stg.tasks)
	}
}

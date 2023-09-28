package optimal_freelancing

import (
	"strconv"
	"testing"
)

var stages = []struct {
	jobs   []map[string]int
	expect int
}{
	{
		jobs:   []map[string]int{},
		expect: 0,
	},
	{
		jobs:   []map[string]int{{"deadline": 1, "payment": 1}},
		expect: 1,
	},
	{
		jobs:   []map[string]int{{"deadline": 1, "payment": 2}, {"deadline": 1, "payment": 1}},
		expect: 2,
	},
	{
		jobs:   []map[string]int{{"deadline": 1, "payment": 1}, {"deadline": 1, "payment": 2}},
		expect: 2,
	},
	{
		jobs:   []map[string]int{{"deadline": 1, "payment": 1}, {"deadline": 2, "payment": 1}},
		expect: 2,
	},
	{
		jobs: []map[string]int{
			{"deadline": 1, "payment": 1},
			{"deadline": 2, "payment": 2},
			{"deadline": 2, "payment": 1},
		},
		expect: 3,
	},
	{
		jobs: []map[string]int{
			{"deadline": 8, "payment": 3},
			{"deadline": 1, "payment": 1},
			{"deadline": 1, "payment": 2},
		},
		expect: 5,
	},
	{
		jobs: []map[string]int{
			{"deadline": 2, "payment": 2},
			{"deadline": 4, "payment": 3},
			{"deadline": 5, "payment": 1},
			{"deadline": 7, "payment": 2},
			{"deadline": 3, "payment": 1},
			{"deadline": 3, "payment": 2},
			{"deadline": 1, "payment": 3},
		},
		expect: 13,
	},
	{
		jobs: []map[string]int{
			{"deadline": 2, "payment": 1},
			{"deadline": 2, "payment": 2},
			{"deadline": 2, "payment": 3},
			{"deadline": 2, "payment": 4},
			{"deadline": 2, "payment": 5},
			{"deadline": 2, "payment": 6},
			{"deadline": 2, "payment": 7},
		},
		expect: 13,
	},
	{
		jobs: []map[string]int{
			{"deadline": 8, "payment": 1},
			{"deadline": 6, "payment": 4},
			{"deadline": 1, "payment": 2},
			{"deadline": 1, "payment": 3},
			{"deadline": 2, "payment": 3},
			{"deadline": 5, "payment": 2},
		},
		expect: 13,
	},
	{
		jobs: []map[string]int{
			{"deadline": 2, "payment": 1},
			{"deadline": 1, "payment": 4},
			{"deadline": 3, "payment": 2},
			{"deadline": 1, "payment": 3},
			{"deadline": 4, "payment": 3},
			{"deadline": 4, "payment": 2},
			{"deadline": 4, "payment": 1},
			{"deadline": 5, "payment": 4},
			{"deadline": 8, "payment": 1},
		},
		expect: 16,
	},
	{
		jobs: []map[string]int{
			{"deadline": 10, "payment": 1},
		},
		expect: 1,
	},
	{
		jobs: []map[string]int{
			{"deadline": 1, "payment": 1},
			{"deadline": 2, "payment": 1},
			{"deadline": 3, "payment": 1},
			{"deadline": 4, "payment": 1},
			{"deadline": 5, "payment": 1},
			{"deadline": 6, "payment": 1},
			{"deadline": 7, "payment": 1},
			{"deadline": 8, "payment": 1},
			{"deadline": 9, "payment": 1},
			{"deadline": 10, "payment": 1},
		},
		expect: 7,
	},
}

func TestOptimalFreelancing(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := OptimalFreelancing(stage.jobs)
			if r != stage.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkOptimalFreelancing(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stage := &stages[i%len(stages)]
		OptimalFreelancing(stage.jobs)
	}
}

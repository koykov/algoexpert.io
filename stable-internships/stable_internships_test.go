package stable_internships

import (
	"fmt"
	"strconv"
	"testing"
)

var stages = []struct {
	interns, teams, expect [][]int
}{
	{
		interns: [][]int{
			{1, 0},
			{0, 1},
		},
		teams: [][]int{
			{0, 1},
			{1, 0},
		},
		expect: [][]int{
			{1, 0},
			{0, 1},
		},
	},
	{
		interns: [][]int{
			{0, 1, 2},
			{0, 2, 1},
			{1, 2, 0},
		},
		teams: [][]int{
			{2, 1, 0},
			{0, 1, 2},
			{0, 2, 1},
		},
		expect: [][]int{
			{1, 0},
			{0, 1},
			{2, 2},
		},
	},
}

func TestStableInternships(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := StableInternships(stg.interns, stg.teams)
			if fmt.Sprintf("%v", r) != fmt.Sprintf("%v", stg.expect) {
				println(fmt.Sprintf("%v", r))
				t.FailNow()
			}
		})
	}
}

func BenchmarkStableInternships(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		StableInternships(stg.interns, stg.teams)
	}
}

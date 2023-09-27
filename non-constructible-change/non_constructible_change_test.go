package non_constructible_change

import (
	"strconv"
	"testing"
)

var stages = []struct {
	coins  []int
	expect int
}{
	{
		coins:  []int{5, 7, 1, 1, 2, 3, 22},
		expect: 20,
	},
	{
		coins:  []int{1, 1, 1, 1, 1},
		expect: 6,
	},
	{
		coins:  []int{1, 5, 1, 1, 1, 10, 15, 20, 100},
		expect: 55,
	},
	{
		coins:  []int{6, 4, 5, 1, 1, 8, 9},
		expect: 3,
	},
	{
		coins:  []int{},
		expect: 1,
	},
	{
		coins:  []int{87},
		expect: 1,
	},
	{
		coins:  []int{5, 6, 1, 1, 2, 3, 4, 9},
		expect: 32,
	},
	{
		coins:  []int{5, 6, 1, 1, 2, 3, 43},
		expect: 19,
	},
	{
		coins:  []int{1, 1},
		expect: 3,
	},
	{
		coins:  []int{2},
		expect: 1,
	},
	{
		coins:  []int{1},
		expect: 2,
	},
	{
		coins:  []int{109, 2000, 8765, 19, 18, 17, 16, 8, 1, 1, 2, 4},
		expect: 87,
	},
	{
		coins:  []int{1, 2, 3, 4, 5, 6, 7},
		expect: 29,
	},
}

func TestNonConstructibleChange(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := NonConstructibleChange(stage.coins)
			if r != stage.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkNonConstructibleChange(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stage := &stages[i%len(stages)]
		NonConstructibleChange(stage.coins)
	}
}

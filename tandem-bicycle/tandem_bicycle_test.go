package tandem_bicycle

import (
	"strconv"
	"testing"
)

var stages = []struct {
	red, blue []int
	fastest   bool
	expect    int
}{
	{
		red:     []int{5, 5, 3, 9, 2},
		blue:    []int{3, 6, 7, 2, 1},
		fastest: true,
		expect:  32,
	},
	{
		red:     []int{5, 5, 3, 9, 2},
		blue:    []int{3, 6, 7, 2, 1},
		fastest: false,
		expect:  25,
	},
	{
		red:     []int{1, 2, 1, 9, 12, 3},
		blue:    []int{3, 3, 4, 6, 1, 2},
		fastest: false,
		expect:  30,
	},
	{
		red:     []int{1, 2, 1, 9, 12, 3},
		blue:    []int{3, 3, 4, 6, 1, 2},
		fastest: true,
		expect:  37,
	},
	{
		red:     []int{1, 2, 1, 9, 12, 3, 1, 54, 21, 231, 32, 1},
		blue:    []int{3, 3, 4, 6, 1, 2, 5, 6, 34, 256, 123, 32},
		fastest: true,
		expect:  816,
	},
	{
		red:     []int{1, 2, 1, 9, 12, 3, 1, 54, 21, 231, 32, 1},
		blue:    []int{3, 3, 4, 6, 1, 2, 5, 6, 34, 256, 123, 32},
		fastest: false,
		expect:  484,
	},
	{
		red:     []int{1},
		blue:    []int{5},
		fastest: true,
		expect:  5,
	},
	{
		red:     []int{1},
		blue:    []int{5},
		fastest: false,
		expect:  5,
	},
	{
		red:     []int{1, 1, 1, 1},
		blue:    []int{1, 1, 1, 1},
		fastest: true,
		expect:  4,
	},
	{
		red:     []int{1, 1, 1, 1},
		blue:    []int{1, 1, 1, 1},
		fastest: false,
		expect:  4,
	},
	{
		red:     []int{1, 1, 1, 1, 2, 2, 2, 2, 9, 11},
		blue:    []int{1, 1, 1, 1, 3, 3, 3, 3, 5, 7},
		fastest: true,
		expect:  48,
	},
	{
		red:     []int{1, 1, 1, 1, 2, 2, 2, 2, 9, 11},
		blue:    []int{1, 1, 1, 1, 3, 3, 3, 3, 5, 7},
		fastest: false,
		expect:  36,
	},
	{
		red:     []int{9, 8, 2, 2, 3, 5, 6},
		blue:    []int{3, 4, 4, 1, 1, 8, 9},
		fastest: true,
		expect:  49,
	},
	{
		red:     []int{9, 8, 2, 2, 3, 5, 6},
		blue:    []int{3, 4, 4, 1, 1, 8, 9},
		fastest: false,
		expect:  35,
	},
	{
		red:     []int{5, 4, 3, 2, 1},
		blue:    []int{1, 2, 3, 4, 5},
		fastest: false,
		expect:  15,
	},
	{
		red:     []int{5, 4, 3, 2, 1},
		blue:    []int{1, 2, 3, 4, 5},
		fastest: true,
		expect:  21,
	},
	{
		red:     []int{5, 4, 3, 2, 1},
		blue:    []int{1, 2, 3, 4, 5},
		fastest: false,
		expect:  15,
	},
}

func TestTandemBicycle(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := TandemBicycle(stage.red, stage.blue, stage.fastest)
			if r != stage.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkTandemBicycle(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stage := &stages[i%len(stages)]
		TandemBicycle(stage.red, stage.blue, stage.fastest)
	}
}

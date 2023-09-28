package class_photos

import (
	"strconv"
	"testing"
)

var stages = []struct {
	red, blue []int
	expect    bool
}{
	{
		red:    []int{5, 8, 1, 3, 4},
		blue:   []int{6, 9, 2, 4, 5},
		expect: true,
	},
	{
		red:    []int{6, 9, 2, 4, 5},
		blue:   []int{5, 8, 1, 3, 4},
		expect: true,
	},
	{
		red:    []int{6, 9, 2, 4, 5, 1},
		blue:   []int{5, 8, 1, 3, 4, 9},
		expect: false,
	},
	{
		red:    []int{6},
		blue:   []int{6},
		expect: false,
	},
	{
		red:    []int{126},
		blue:   []int{125},
		expect: true,
	},
	{
		red:    []int{1, 2, 3, 4, 5},
		blue:   []int{2, 3, 4, 5, 6},
		expect: true,
	},
	{
		red:    []int{1, 1, 1, 1, 1, 1, 1, 1},
		blue:   []int{5, 6, 7, 2, 3, 1, 2, 3},
		expect: false,
	},
	{
		red:    []int{1, 1, 1, 1, 1, 1, 1, 1},
		blue:   []int{2, 2, 2, 2, 2, 2, 2, 2},
		expect: true,
	},
	{
		red:    []int{125},
		blue:   []int{126},
		expect: true,
	},
	{
		red:    []int{19, 2, 4, 6, 2, 3, 1, 1, 4},
		blue:   []int{21, 5, 4, 4, 4, 4, 4, 4, 4},
		expect: false,
	},
	{
		red:    []int{19, 19, 21, 1, 1, 1, 1, 1},
		blue:   []int{20, 5, 4, 4, 4, 4, 4, 4},
		expect: false,
	},
	{
		red:    []int{3, 5, 6, 8, 2},
		blue:   []int{2, 4, 7, 5, 1},
		expect: true,
	},
	{
		red:    []int{3, 5, 6, 8, 2, 1},
		blue:   []int{2, 4, 7, 5, 1, 6},
		expect: false,
	},
	{
		red:    []int{4, 5},
		blue:   []int{5, 4},
		expect: false,
	},
	{
		red:    []int{5, 6},
		blue:   []int{5, 4},
		expect: true,
	},
}

func TestClassPhotos(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := ClassPhotos(stage.red, stage.blue)
			if r != stage.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkClassPhotos(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stage := &stages[i%len(stages)]
		ClassPhotos(stage.red, stage.blue)
	}
}

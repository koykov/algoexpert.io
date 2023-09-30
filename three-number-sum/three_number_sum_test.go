package three_number_sum

import (
	"fmt"
	"strconv"
	"testing"
)

var stages = []struct {
	array  []int
	target int
	expect [][]int
}{
	{
		array:  []int{12, 3, 1, 2, -6, 5, -8, 6},
		target: 0,
		expect: [][]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 1, 5}},
	},
	{
		array:  []int{1, 2, 3},
		target: 6,
		expect: [][]int{{1, 2, 3}},
	},
	{
		array:  []int{1, 2, 3},
		target: 7,
		expect: [][]int{},
	},
	{
		array:  []int{8, 10, -2, 49, 14},
		target: 57,
		expect: [][]int{{-2, 10, 49}},
	},
	{
		array:  []int{12, 3, 1, 2, -6, 5, 0, -8, -1},
		target: 0,
		expect: [][]int{{-8, 3, 5}, {-6, 1, 5}, {-1, 0, 1}},
	},
	{
		array:  []int{12, 3, 1, 2, -6, 5, 0, -8, -1, 6},
		target: 0,
		expect: [][]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 0, 6}, {-6, 1, 5}, {-1, 0, 1}},
	},
	{
		array:  []int{12, 3, 1, 2, -6, 5, 0, -8, -1, 6, -5},
		target: 0,
		expect: [][]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 0, 6}, {-6, 1, 5}, {-5, -1, 6}, {-5, 0, 5}, {-5, 2, 3}, {-1, 0, 1}},
	},
	{
		array:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15},
		target: 18,
		expect: [][]int{{1, 2, 15}, {1, 8, 9}, {2, 7, 9}, {3, 6, 9}, {3, 7, 8}, {4, 5, 9}, {4, 6, 8}, {5, 6, 7}},
	},
	{
		array:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15},
		target: 32,
		expect: [][]int{{8, 9, 15}},
	},
	{
		array:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15},
		target: 33,
		expect: [][]int{},
	},
	{
		array:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15},
		target: 5,
		expect: [][]int{},
	},
}

func TestThreeNumberSum(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := ThreeNumberSum(stg.array, stg.target)
			if rs, es := fmt.Sprintf("%#v", r), fmt.Sprintf("%#v", stg.expect); rs != es {
				t.FailNow()
			}
		})
	}
}

func BenchmarkThreeNumberSum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		ThreeNumberSum(stg.array, stg.target)
	}
}

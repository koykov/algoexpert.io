package move_element_to_end

import (
	"fmt"
	"strconv"
	"testing"
)

var stages = []struct {
	arr    []int
	num    int
	expect []int
}{
	{
		arr:    []int{2, 1, 2, 2, 2, 3, 4, 2},
		num:    2,
		expect: []int{1, 3, 4, 2, 2, 2, 2, 2},
	},
	{
		arr:    []int{},
		num:    3,
		expect: []int{},
	},
	{
		arr:    []int{1, 2, 4, 5, 6},
		num:    3,
		expect: []int{1, 2, 4, 5, 6},
	},
	{
		arr:    []int{3, 3, 3, 3, 3},
		num:    3,
		expect: []int{3, 3, 3, 3, 3},
	},
	{
		arr:    []int{3, 1, 2, 4, 5},
		num:    3,
		expect: []int{1, 2, 4, 5, 3},
	},
	{
		arr:    []int{1, 2, 4, 5, 3},
		num:    3,
		expect: []int{1, 2, 4, 5, 3},
	},
	{
		arr:    []int{1, 2, 3, 4, 5},
		num:    3,
		expect: []int{1, 2, 4, 5, 3},
	},
	{
		arr:    []int{2, 4, 2, 5, 6, 2, 2},
		num:    2,
		expect: []int{4, 5, 6, 2, 2, 2, 2},
	},
	{
		arr:    []int{5, 5, 5, 5, 5, 5, 1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12},
		num:    5,
		expect: []int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 5, 5, 5, 5, 5, 5},
	},
	{
		arr:    []int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 5, 5, 5, 5, 5, 5},
		num:    5,
		expect: []int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 5, 5, 5, 5, 5, 5},
	},
	{
		arr:    []int{5, 1, 2, 5, 5, 3, 4, 6, 7, 5, 8, 9, 10, 11, 5, 5, 12},
		num:    5,
		expect: []int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 5, 5, 5, 5, 5, 5},
	},
}

func TestMoveElementToEnd(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			stg := &stages[i]
			r := MoveElementToEnd(stg.arr, stg.num)
			if fmt.Sprintf("%#v", r) != fmt.Sprintf("%#v", stg.expect) {
				t.FailNow()
			}
		})
	}
}

func BenchmarkMoveElementToEnd(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		MoveElementToEnd(stg.arr, stg.num)
	}
}

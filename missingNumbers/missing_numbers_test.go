package missingNumbers

import (
	"fmt"
	"strconv"
	"testing"
)

var stages = []struct {
	arr, expect []int
}{
	{
		arr:    []int{},
		expect: []int{1, 2},
	},
	{
		arr:    []int{1},
		expect: []int{2, 3},
	},
	{
		arr:    []int{2},
		expect: []int{1, 3},
	},
	{
		arr:    []int{3},
		expect: []int{1, 2},
	},
	{
		arr:    []int{1, 3},
		expect: []int{2, 4},
	},
	{
		arr:    []int{3, 1},
		expect: []int{2, 4},
	},
	{
		arr:    []int{1, 2, 3},
		expect: []int{4, 5},
	},
	{
		arr:    []int{3, 2, 1},
		expect: []int{4, 5},
	},
	{
		arr:    []int{3, 1, 2},
		expect: []int{4, 5},
	},
	{
		arr:    []int{3, 4, 5},
		expect: []int{1, 2},
	},
	{
		arr:    []int{4, 5, 3},
		expect: []int{1, 2},
	},
	{
		arr:    []int{1, 3, 4, 5},
		expect: []int{2, 6},
	},
	{
		arr:    []int{4, 5, 1, 3},
		expect: []int{2, 6},
	},
	{
		arr:    []int{1, 2, 4, 5, 7},
		expect: []int{3, 6},
	},
	{
		arr:    []int{1, 2, 7, 5, 4},
		expect: []int{3, 6},
	},
	{
		arr:    []int{1, 2, 3, 4, 5, 6, 7},
		expect: []int{8, 9},
	},
	{
		arr:    []int{7, 6, 5, 4, 3, 2, 1},
		expect: []int{8, 9},
	},
	{
		arr:    []int{3, 5, 1, 2, 4, 7, 6},
		expect: []int{8, 9},
	},
	{
		arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18, 20, 21, 22},
		expect: []int{14, 19},
	},
	{
		arr:    []int{3, 5, 7, 8, 1, 10, 11, 2, 4, 13, 17, 22, 18, 21, 16, 20, 6, 9, 15, 12},
		expect: []int{14, 19},
	},
	{
		arr:    []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22},
		expect: []int{1, 2},
	},
	{
		arr:    []int{14, 15, 16, 17, 18, 19, 20, 21, 22, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
		expect: []int{1, 2},
	},
	{
		arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22},
		expect: []int{23, 24},
	},
	{
		arr:    []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		expect: []int{23, 24},
	},
}

func TestMissingNumbers(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			stg := &stages[i]
			r := MissingNumbers(stg.arr)
			if fmt.Sprintf("%#v", r) != fmt.Sprintf("%#v", stg.expect) {
				t.FailNow()
			}
		})
	}
}

func BenchmarkMissingNumbers(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		MissingNumbers(stg.arr)
	}
}

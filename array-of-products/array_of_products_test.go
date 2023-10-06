package array_of_products

import (
	"fmt"
	"strconv"
	"testing"
)

var stages = []struct {
	arr, expect []int
}{
	{
		arr:    []int{5, 1, 4, 2},
		expect: []int{8, 40, 10, 20},
	},
	{
		arr:    []int{1, 8, 6, 2, 4},
		expect: []int{384, 48, 64, 192, 96},
	},
	{
		arr:    []int{-5, 2, -4, 14, -6},
		expect: []int{672, -1680, 840, -240, 560},
	},
	{
		arr:    []int{9, 3, 2, 1, 9, 5, 3, 2},
		expect: []int{1620, 4860, 7290, 14580, 1620, 2916, 4860, 7290},
	},
	{
		arr:    []int{4, 4},
		expect: []int{4, 4},
	},
	{
		arr:    []int{0, 0, 0, 0},
		expect: []int{0, 0, 0, 0},
	},
	{
		arr:    []int{1, 1, 1, 1},
		expect: []int{1, 1, 1, 1},
	},
	{
		arr:    []int{-1, -1, -1},
		expect: []int{1, 1, 1},
	},
	{
		arr:    []int{-1, -1, -1, -1},
		expect: []int{-1, -1, -1, -1},
	},
	{
		arr:    []int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		expect: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		arr:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		expect: []int{362880, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
}

func TestArrayOfProducts(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			stg := &stages[i]
			r := ArrayOfProducts(stg.arr)
			if fmt.Sprintf("%#v", r) != fmt.Sprintf("%#v", stg.expect) {
				t.FailNow()
			}
		})
	}
}

func BenchmarkArrayOfProducts(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		ArrayOfProducts(stg.arr)
	}
}

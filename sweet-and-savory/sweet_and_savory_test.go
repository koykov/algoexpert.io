package sweet_and_savory

import (
	"fmt"
	"strconv"
	"testing"
)

var stages = []struct {
	arr    []int
	target int
	expect []int
}{
	{
		arr:    []int{-3, -5, 1, 7},
		target: 8,
		expect: []int{-3, 7},
	},
	{
		arr:    []int{-5, 10},
		target: 4,
		expect: []int{0, 0},
	},
}

func TestSweetAndSavory(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			stg := &stages[i]
			r := SweetAndSavory(stg.arr, stg.target)
			if fmt.Sprintf("%#v", r) != fmt.Sprintf("%#v", stg.expect) {
				t.FailNow()
			}
		})
	}
}

func BenchmarkSweetAndSavory(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		SweetAndSavory(stg.arr, stg.target)
	}
}

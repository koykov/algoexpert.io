package majority_element

import (
	"strconv"
	"testing"
)

var stages = []struct {
	arr    []int
	expect int
}{
	{
		arr:    []int{2},
		expect: 2,
	},
	{
		arr:    []int{1, 2, 1},
		expect: 1,
	},
	{
		arr:    []int{3, 3, 1},
		expect: 3,
	},
	{
		arr:    []int{4, 5, 5},
		expect: 5,
	},
	{
		arr:    []int{1, 2, 3, 2, 2, 1, 2},
		expect: 2,
	},
	{
		arr:    []int{1, 2, 3, 2, 3, 2, 2, 4, 2},
		expect: 2,
	},
	{
		arr:    []int{1, 1, 1, 1, 1, 1, 1},
		expect: 1,
	},
	{
		arr:    []int{5, 4, 3, 2, 1, 1, 1, 1, 1},
		expect: 1,
	},
	{
		arr:    []int{1, 1, 1, 1, 1, 5, 4, 3, 2},
		expect: 1,
	},
	{
		arr:    []int{1, 1, 1, 1, 2, 2, 2, 2, 2},
		expect: 2,
	},
	{
		arr:    []int{435, 6543, 6543, 435, 535, 6543, 6543, 12, 43, 6543, 6543},
		expect: 6543,
	},
	{
		arr:    []int{1, 2, 2, 2, 1},
		expect: 2,
	},
	{
		arr:    []int{1, 2, 3, 4, 4, 4, 4, 4, 4, 4, 3, 2, 1},
		expect: 4,
	},
	{
		arr:    []int{1, 2, 3, 2, 2, 4, 2, 2, 5, 2, 1},
		expect: 2,
	},
}

func TestMajorityElement(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			stg := &stages[i]
			r := MajorityElement(stg.arr)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkMajorityElement(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		MajorityElement(stg.arr)
	}
}

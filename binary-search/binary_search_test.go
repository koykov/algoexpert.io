package binary_search

import (
	"strconv"
	"testing"
)

var stages = []struct {
	array  []int
	target int
	expect int
}{
	{
		array:  []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73},
		target: 33,
		expect: 3,
	},
	{
		array:  []int{1, 5, 23, 111},
		target: 111,
		expect: 3,
	},
	{
		array:  []int{1, 5, 23, 111},
		target: 5,
		expect: 1,
	},
	{
		array:  []int{1, 5, 23, 111},
		target: 35,
		expect: -1,
	},
	{
		array:  []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73},
		target: 0,
		expect: 0,
	},
	{
		array:  []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73},
		target: 1,
		expect: 1,
	},
	{
		array:  []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73},
		target: 21,
		expect: 2,
	},
	{
		array:  []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73},
		target: 45,
		expect: 4,
	},
	{
		array:  []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73},
		target: 61,
		expect: 6,
	},
	{
		array:  []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73},
		target: 71,
		expect: 7,
	},
	{
		array:  []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73},
		target: 72,
		expect: 8,
	},
	{
		array:  []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73},
		target: 73,
		expect: 9,
	},
	{
		array:  []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73},
		target: 70,
		expect: -1,
	},
	{
		array:  []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73, 355},
		target: 355,
		expect: 10,
	},
	{
		array:  []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73, 355},
		target: 354,
		expect: -1,
	},
	{
		array:  []int{1, 5, 23, 111},
		target: 120,
		expect: -1,
	},
	{
		array:  []int{5, 23, 111},
		target: 3,
		expect: -1,
	},
}

func TestBinarySearch(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := BinarySearch(stg.array, stg.target)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		BinarySearch(stg.array, stg.target)
	}
}

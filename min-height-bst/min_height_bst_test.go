package min_height_bst

import (
	"strconv"
	"testing"
)

var stages = []struct {
	arr    []int
	expect string
}{
	{
		arr:    []int{1, 2, 5, 7, 10, 13, 14, 15, 22},
		expect: "{value:10,left:{value:5,left:{value:2,left:{value:1}},right:{value:7}},right:{value:15,left:{value:14,left:{value:13}},right:{value:22}}}",
	},
	{
		arr:    []int{1},
		expect: "{value:1}",
	},
	{
		arr:    []int{1, 2},
		expect: "{value:2,left:{value:1}}",
	},
	{
		arr:    []int{1, 2, 5},
		expect: "{value:2,left:{value:1},right:{value:5}}",
	},
	{
		arr:    []int{1, 2, 5, 7},
		expect: "{value:5,left:{value:2,left:{value:1}},right:{value:7}}",
	},
	{
		arr:    []int{1, 2, 5, 7, 10},
		expect: "{value:5,left:{value:2,left:{value:1}},right:{value:10,left:{value:7}}}",
	},
	{
		arr:    []int{1, 2, 5, 7, 10, 13},
		expect: "{value:7,left:{value:2,left:{value:1},right:{value:5}},right:{value:13,left:{value:10}}}",
	},
	{
		arr:    []int{1, 2, 5, 7, 10, 13, 14},
		expect: "{value:7,left:{value:2,left:{value:1},right:{value:5}},right:{value:13,left:{value:10},right:{value:14}}}",
	},
	{
		arr:    []int{1, 2, 5, 7, 10, 13, 14, 15},
		expect: "{value:10,left:{value:5,left:{value:2,left:{value:1}},right:{value:7}},right:{value:14,left:{value:13},right:{value:15}}}",
	},
	{
		arr:    []int{1, 2, 5, 7, 10, 13, 14, 15, 22},
		expect: "{value:10,left:{value:5,left:{value:2,left:{value:1}},right:{value:7}},right:{value:15,left:{value:14,left:{value:13}},right:{value:22}}}",
	},
	{
		arr:    []int{1, 2, 5, 7, 10, 13, 14, 15, 22, 28},
		expect: "{value:13,left:{value:5,left:{value:2,left:{value:1}},right:{value:10,left:{value:7}}},right:{value:22,left:{value:15,left:{value:14}},right:{value:28}}}",
	},
	{
		arr:    []int{1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32},
		expect: "{value:13,left:{value:5,left:{value:2,left:{value:1}},right:{value:10,left:{value:7}}},right:{value:22,left:{value:15,left:{value:14}},right:{value:32,left:{value:28}}}}",
	},
	{
		arr:    []int{1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36},
		expect: "{value:14,left:{value:7,left:{value:2,left:{value:1},right:{value:5}},right:{value:13,left:{value:10}}},right:{value:28,left:{value:22,left:{value:15}},right:{value:36,left:{value:32}}}}",
	},
	{
		arr:    []int{1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36, 89},
		expect: "{value:14,left:{value:7,left:{value:2,left:{value:1},right:{value:5}},right:{value:13,left:{value:10}}},right:{value:32,left:{value:22,left:{value:15},right:{value:28}},right:{value:89,left:{value:36}}}}",
	},
	{
		arr:    []int{1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36, 89, 92},
		expect: "{value:15,left:{value:7,left:{value:2,left:{value:1},right:{value:5}},right:{value:13,left:{value:10},right:{value:14}}},right:{value:36,left:{value:28,left:{value:22},right:{value:32}},right:{value:92,left:{value:89}}}}",
	},
	{
		arr:    []int{1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36, 89, 92, 9000},
		expect: "{value:15,left:{value:7,left:{value:2,left:{value:1},right:{value:5}},right:{value:13,left:{value:10},right:{value:14}}},right:{value:36,left:{value:28,left:{value:22},right:{value:32}},right:{value:92,left:{value:89},right:{value:9000}}}}",
	},
	{
		arr:    []int{1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36, 89, 92, 9000, 9001},
		expect: "{value:22,left:{value:10,left:{value:5,left:{value:2,left:{value:1}},right:{value:7}},right:{value:14,left:{value:13},right:{value:15}}},right:{value:89,left:{value:32,left:{value:28},right:{value:36}},right:{value:9000,left:{value:92},right:{value:9001}}}}",
	},
}

func TestBST(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tree := MinHeightBST(stg.arr)
			r := tree.Marshal()
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkBST(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		MinHeightBST(stg.arr)
	}
}

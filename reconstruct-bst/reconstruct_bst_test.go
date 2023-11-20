package reconstruct_bst

import (
	"strconv"
	"testing"
)

var stages = []struct {
	arr    []int
	expect string
}{
	{
		arr:    []int{10, 4, 2, 1, 5, 17, 19, 18},
		expect: "{value:10,left:{value:4,left:{value:2,left:{value:1}},right:{value:5}},right:{value:17,right:{value:19,left:{value:18}}}}",
	},
	{
		arr:    []int{100},
		expect: "{value:100}",
	},
	{
		arr:    []int{10, 9, 8, 7, 6, 5},
		expect: "{value:10,left:{value:9,left:{value:8,left:{value:7,left:{value:6,left:{value:5}}}}}}",
	},
	{
		arr:    []int{5, 6, 7, 8},
		expect: "{value:5,right:{value:6,right:{value:7,right:{value:8}}}}",
	},
	{
		arr:    []int{5, -10, -5, 6, 9, 7},
		expect: "{value:5,left:{value:-10,right:{value:-5}},right:{value:6,right:{value:9,left:{value:7}}}}",
	},
	{
		arr:    []int{10, 4, 2, 1, 3, 5, 6, 9, 7, 17, 19, 18},
		expect: "{value:10,left:{value:4,left:{value:2,left:{value:1},right:{value:3}},right:{value:5,right:{value:6,right:{value:9,left:{value:7}}}}},right:{value:17,right:{value:19,left:{value:18}}}}",
	},
	{
		arr:    []int{1, 0, 2},
		expect: "{value:1,left:{value:0},right:{value:2}}",
	},
	{
		arr:    []int{2, 0, 1},
		expect: "{value:2,left:{value:0,right:{value:1}}}",
	},
	{
		arr:    []int{2, 0, 1, 4, 3},
		expect: "{value:2,left:{value:0,right:{value:1}},right:{value:4,left:{value:3}}}",
	},
	{
		arr:    []int{2, 0, 1, 4, 3, 3},
		expect: "{value:2,left:{value:0,right:{value:1}},right:{value:4,left:{value:3,right:{value:3}}}}",
	},
	{
		arr:    []int{2, 0, 1, 3, 4, 3},
		expect: "{value:2,left:{value:0,right:{value:1}},right:{value:3,right:{value:4,left:{value:3}}}}",
	},
	{
		arr:    []int{10, 4, 2, 1, 3, 5, 5, 6, 5, 5, 9, 7, 17, 19, 18, 18, 19},
		expect: "{value:10,left:{value:4,left:{value:2,left:{value:1},right:{value:3}},right:{value:5,right:{value:5,right:{value:6,left:{value:5,right:{value:5}},right:{value:9,left:{value:7}}}}}},right:{value:17,right:{value:19,left:{value:18,right:{value:18}},right:{value:19}}}}",
	},
}

func TestBST(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tree := ReconstructBst(stg.arr)
			r := tree.Marshal()
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stg := &stages[i]
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			b.ReportAllocs()
			ReconstructBst(stg.arr)
		})
	}
}

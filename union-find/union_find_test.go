package union_find

import (
	"strconv"
	"testing"
)

type scall struct {
	method string
	args   []int
	expect int
}

var stages = []struct {
	calls []scall
}{
	// {
	// 	calls: []scall{
	// 		{
	// 			args:   []int{0},
	// 			method: "find",
	// 		},
	// 		{
	// 			args:   []int{1},
	// 			method: "find",
	// 		},
	// 	},
	// },
	// {
	// 	calls: []scall{
	// 		{
	// 			args:   []int{0},
	// 			method: "find",
	// 		},
	// 		{
	// 			args:   []int{1},
	// 			method: "find",
	// 		},
	// 		{
	// 			args:   []int{0, 1},
	// 			method: "union",
	// 		},
	// 		{
	// 			args:   []int{0},
	// 			method: "find",
	// 		},
	// 		{
	// 			args:   []int{1},
	// 			method: "find",
	// 		},
	// 	},
	// },
	// {
	// 	calls: []scall{
	// 		{
	// 			args:   []int{0},
	// 			method: "createSet",
	// 		},
	// 		{
	// 			args:   []int{1},
	// 			method: "createSet",
	// 		},
	// 		{
	// 			args:   []int{0},
	// 			method: "find",
	// 		},
	// 		{
	// 			args:   []int{1},
	// 			method: "find",
	// 		},
	// 	},
	// },
	{
		calls: []scall{
			{
				args:   []int{10},
				method: "createSet",
			},
			{
				args:   []int{5},
				method: "createSet",
			},
			{
				args:   []int{10},
				method: "find",
			},
			{
				args:   []int{5},
				method: "find",
			},
			{
				args:   []int{10, 5},
				method: "union",
			},
			{
				args:   []int{10},
				method: "find",
			},
			{
				args:   []int{5},
				method: "find",
			},
		},
	},
}

func TestStableInternships(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			dsu := NewUnionFind()
			for j := 0; j < len(stg.calls); j++ {
				switch stg.calls[j].method {
				case "createSet":
					dsu.CreateSet(stg.calls[j].args[0])
				case "union":
					dsu.Union(stg.calls[j].args[0], stg.calls[j].args[1])
				case "find":
					x := dsu.Find(stg.calls[j].args[0])
					println(*x)
					if *x != stg.calls[j].expect {
						// t.FailNow()
					}
				}
			}
		})
	}
}

func BenchmarkStableInternships(b *testing.B) {
	// b.ReportAllocs()
	// for i := 0; i < b.N; i++ {
	// 	stg := &stages[i%len(stages)]
	// 	StableInternships(stg.interns, stg.teams)
	// }
}

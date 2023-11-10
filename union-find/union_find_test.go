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
	{
		calls: []scall{
			{
				args:   []int{0},
				method: "find",
			},
			{
				args:   []int{1},
				method: "find",
			},
		},
	},
	{
		calls: []scall{
			{
				args:   []int{0},
				method: "find",
			},
			{
				args:   []int{1},
				method: "find",
			},
			{
				args:   []int{0, 1},
				method: "union",
			},
			{
				args:   []int{0},
				method: "find",
			},
			{
				args:   []int{1},
				method: "find",
			},
		},
	},
	{
		calls: []scall{
			{
				args:   []int{0},
				method: "createSet",
			},
			{
				args:   []int{1},
				method: "createSet",
			},
			{
				args:   []int{0},
				method: "find",
				expect: 0,
			},
			{
				args:   []int{1},
				method: "find",
				expect: 1,
			},
		},
	},
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
				expect: 10,
			},
			{
				args:   []int{5},
				method: "find",
				expect: 5,
			},
			{
				args:   []int{10, 5},
				method: "union",
			},
			{
				args:   []int{10},
				method: "find",
				expect: 10,
			},
			{
				args:   []int{5},
				method: "find",
				expect: 10,
			},
		},
	},
	{
		calls: []scall{
			{
				args:   []int{0},
				method: "createSet",
			},
			{
				args:   []int{2},
				method: "createSet",
			},
			{
				args:   []int{0, 2},
				method: "union",
			},
			{
				args:   []int{3},
				method: "createSet",
			},
			{
				args:   []int{1},
				method: "createSet",
			},
			{
				args:   []int{1, 3},
				method: "union",
			},
			{
				args:   []int{0},
				method: "find",
				expect: 0,
			},
			{
				args:   []int{1},
				method: "find",
				expect: 1,
			},
			{
				args:   []int{2},
				method: "find",
				expect: 0,
			},
			{
				args:   []int{3},
				method: "find",
				expect: 1,
			},
			{
				args:   []int{3, 0},
				method: "union",
			},
			{
				args:   []int{0},
				method: "find",
				expect: 1,
			},
			{
				args:   []int{1},
				method: "find",
				expect: 1,
			},
			{
				args:   []int{2},
				method: "find",
				expect: 1,
			},
			{
				args:   []int{3},
				method: "find",
				expect: 1,
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
					if x == nil {
						continue
					}
					if *x != stg.calls[j].expect {
						t.FailNow()
					}
				}
			}
		})
	}
}

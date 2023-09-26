package shift_linked_list

import (
	"strconv"
	"testing"
)

type node struct {
	id    string
	next  string
	value int
}

type stage struct {
	nodes []node
	exp   string
	k     int
}

var stages = []stage{
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "4 -> 5 -> 0 -> 1 -> 2 -> 3",
		k:   2,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "0 -> 1 -> 2 -> 3 -> 4 -> 5",
		k:   0,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "5 -> 0 -> 1 -> 2 -> 3 -> 4",
		k:   1,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "3 -> 4 -> 5 -> 0 -> 1 -> 2",
		k:   3,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "2 -> 3 -> 4 -> 5 -> 0 -> 1",
		k:   4,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "1 -> 2 -> 3 -> 4 -> 5 -> 0",
		k:   5,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "0 -> 1 -> 2 -> 3 -> 4 -> 5",
		k:   6,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "4 -> 5 -> 0 -> 1 -> 2 -> 3",
		k:   8,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "4 -> 5 -> 0 -> 1 -> 2 -> 3",
		k:   14,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "0 -> 1 -> 2 -> 3 -> 4 -> 5",
		k:   18,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "1 -> 2 -> 3 -> 4 -> 5 -> 0",
		k:   -1,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "2 -> 3 -> 4 -> 5 -> 0 -> 1",
		k:   -2,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "3 -> 4 -> 5 -> 0 -> 1 -> 2",
		k:   -3,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "4 -> 5 -> 0 -> 1 -> 2 -> 3",
		k:   -4,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "5 -> 0 -> 1 -> 2 -> 3 -> 4",
		k:   -5,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "0 -> 1 -> 2 -> 3 -> 4 -> 5",
		k:   -6,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "2 -> 3 -> 4 -> 5 -> 0 -> 1",
		k:   -8,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "2 -> 3 -> 4 -> 5 -> 0 -> 1",
		k:   -14,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "", value: 5},
		},
		exp: "0 -> 1 -> 2 -> 3 -> 4 -> 5",
		k:   -18,
	},
	{
		nodes: []node{
			{id: "0", next: "1", value: 0},
			{id: "1", next: "4", value: 1},
			{id: "2", next: "", value: 2},
		},
		exp: "1 -> 2 -> 0",
		k:   2,
	},
}

func nodesToLL(nodes []node) *LinkedList {
	head := LinkedList{Value: nodes[0].value}
	node_ := &head
	for i := 1; i < len(nodes); i++ {
		node_.Next = &LinkedList{Value: nodes[i].value}
		node_ = node_.Next
	}
	return &head
}

func ll2bytes(buf []byte, head *LinkedList) []byte {
	node_ := head
	for node_ != nil {
		buf = strconv.AppendInt(buf, int64(node_.Value), 10)
		if node_.Next != nil {
			buf = append(buf, " -> "...)
		}
		node_ = node_.Next
	}
	return buf
}

func TestShiftLinkedList(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ll := nodesToLL(stg.nodes)
			ll1 := ShiftLinkedList(ll, stg.k)
			p := string(ll2bytes(nil, ll1))
			if p != stg.exp {
				t.FailNow()
			}
		})
	}
}

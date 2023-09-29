package middle_node

import (
	"strconv"
	"testing"
)

type node struct {
	id, next string
	value    int
}

var stages = []struct {
	nodes  []node
	head   *LinkedList
	expect string
}{
	{
		nodes: []node{
			{id: "1", next: "", value: 1},
		},
		expect: "1",
	},
	{
		nodes: []node{
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "", value: 3},
		},
		expect: "2 -> 3",
	},
	{
		nodes: []node{
			{id: "5", next: "7", value: 5},
			{id: "7", next: "9", value: 7},
			{id: "9", next: "", value: 9},
		},
		expect: "7 -> 9",
	},
	{
		nodes: []node{
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "", value: 4},
		},
		expect: "3 -> 4",
	},
	{
		nodes: []node{
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "6", value: 5},
			{id: "6", next: "7", value: 6},
			{id: "7", next: "8", value: 7},
			{id: "8", next: "9", value: 8},
			{id: "9", next: "", value: 9},
		},
		expect: "5 -> 6 -> 7 -> 8 -> 9",
	},
	{
		nodes: []node{
			{id: "1", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "6", value: 5},
			{id: "6", next: "7", value: 6},
			{id: "7", next: "8", value: 7},
			{id: "8", next: "9", value: 8},
			{id: "9", next: "10", value: 9},
			{id: "10", next: "", value: 10},
		},
		expect: "6 -> 7 -> 8 -> 9 -> 10",
	},
	{
		nodes: []node{
			{id: "1", next: "3", value: 1},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "9", value: 5},
			{id: "9", next: "10", value: 9},
			{id: "10", next: "", value: 10},
		},
		expect: "5 -> 9 -> 10",
	},
	{
		nodes: []node{
			{id: "1", next: "1-2", value: 1},
			{id: "1-2", next: "1-3", value: 1},
			{id: "1-3", next: "3", value: 1},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "5-2", value: 5},
			{id: "5-2", next: "5-3", value: 5},
			{id: "5-3", next: "5-4", value: 5},
			{id: "5-4", next: "10", value: 5},
			{id: "10", next: "", value: 10},
		},
		expect: "5 -> 5 -> 5 -> 5 -> 10",
	}, {
		nodes: []node{
			{id: "1", next: "2", value: 1},
			{id: "2", next: "1-2", value: 2},
			{id: "1-2", next: "4", value: 1},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "5-2", value: 5},
			{id: "5-2", next: "7", value: 5},
			{id: "7", next: "10", value: 7},
			{id: "10", next: "", value: 10},
		},
		expect: "5 -> 5 -> 7 -> 10",
	},
	{
		nodes: []node{
			{id: "1", next: "1-2", value: 1},
			{id: "1-2", next: "1-3", value: 1},
			{id: "1-3", next: "1-4", value: 1},
			{id: "1-4", next: "2", value: 1},
			{id: "2", next: "3", value: 2},
			{id: "3", next: "4", value: 3},
			{id: "4", next: "5", value: 4},
			{id: "5", next: "5-2", value: 5},
			{id: "5-2", next: "7", value: 5},
			{id: "7", next: "10", value: 7},
			{id: "10", next: "", value: 10},
		},
		expect: "3 -> 4 -> 5 -> 5 -> 7 -> 10",
	},
}

func init() {
	for i := 0; i < len(stages); i++ {
		stages[i].head = nodesToLL(stages[i].nodes)
	}
}

func TestMiddleNode(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := MiddleNode(stg.head)
			p := string(ll2bytes(nil, r))
			if p != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkMiddleNode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MiddleNode(stages[i%len(stages)].head)
	}
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

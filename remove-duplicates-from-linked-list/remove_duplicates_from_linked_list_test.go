package remove_duplicates_from_linked_list

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
			{id: "1", next: "1-2", value: 1},
			{id: "1-2", next: "1-3", value: 1},
			{id: "1-3", next: "2", value: 1},
			{id: "2", next: "3", value: 3},
			{id: "3", next: "3-2", value: 4},
			{id: "3-2", next: "3-3", value: 4},
			{id: "3-3", next: "4", value: 4},
			{id: "4", next: "5", value: 5},
			{id: "5", next: "5-2", value: 6},
			{id: "5-2", next: "", value: 6},
		},
		expect: "1 -> 3 -> 4 -> 5 -> 6",
	},
	{
		nodes: []node{
			{id: "1", next: "1-2", value: 1},
			{id: "1-2", next: "1-3", value: 1},
			{id: "1-3", next: "1-4", value: 1},
			{id: "1-4", next: "1-5", value: 1},
			{id: "1-5", next: "4", value: 1},
			{id: "4", next: "4-2", value: 4},
			{id: "4-2", next: "5", value: 4},
			{id: "5", next: "6", value: 5},
			{id: "6", next: "6-2", value: 6},
			{id: "6-2", next: "", value: 6},
		},
		expect: "1 -> 4 -> 5 -> 6",
	},
	{
		nodes: []node{
			{id: "1", next: "1-2", value: 1},
			{id: "1-2", next: "1-3", value: 1},
			{id: "1-3", next: "1-4", value: 1},
			{id: "1-4", next: "1-5", value: 1},
			{id: "1-5", next: "1-6", value: 1},
			{id: "1-6", next: "1-7", value: 1},
			{id: "1-7", next: "", value: 1},
		},
		expect: "1",
	},
	{
		nodes: []node{
			{id: "1", next: "9", value: 1},
			{id: "9", next: "11", value: 9},
			{id: "11", next: "15", value: 11},
			{id: "15", next: "15-2", value: 15},
			{id: "15-2", next: "16", value: 15},
			{id: "16", next: "17", value: 16},
			{id: "17", next: "", value: 17},
		},
		expect: "1 -> 9 -> 11 -> 15 -> 16 -> 17",
	},
	{
		nodes: []node{
			{id: "1", next: "", value: 1},
		},
		expect: "1",
	},
	{
		nodes: []node{
			{id: "-5", next: "-1", value: -5},
			{id: "-1", next: "-1-2", value: -1},
			{id: "-1-2", next: "-1-3", value: -1},
			{id: "-1-3", next: "5", value: -1},
			{id: "5", next: "5-2", value: 5},
			{id: "5-2", next: "5-3", value: 5},
			{id: "5-3", next: "8", value: 5},
			{id: "8", next: "8-2", value: 8},
			{id: "8-2", next: "9", value: 8},
			{id: "9", next: "10", value: 9},
			{id: "10", next: "11", value: 10},
			{id: "11", next: "11-2", value: 11},
			{id: "11-2", next: "", value: 11},
		},
		expect: "-5 -> -1 -> 5 -> 8 -> 9 -> 10 -> 11",
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
			{id: "10", next: "11", value: 10},
			{id: "11", next: "12", value: 11},
			{id: "12", next: "12-2", value: 12},
			{id: "12-2", next: "", value: 12},
		},
		expect: "1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> 11 -> 12",
	},
}

func init() {
	for i := 0; i < len(stages); i++ {
		stages[i].head = nodesToLL(stages[i].nodes)
	}
}

func TestRemoveDuplicatesFromLinkedList(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := RemoveDuplicatesFromLinkedList(stg.head)
			p := string(ll2bytes(nil, r))
			if p != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkRemoveDuplicatesFromLinkedList(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		RemoveDuplicatesFromLinkedList(stages[i%len(stages)].head)
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

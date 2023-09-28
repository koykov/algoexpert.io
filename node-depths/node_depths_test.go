package node_depths

import (
	"strconv"
	"testing"
)

type node struct {
	id    string
	left  string
	right string
	value int
}

var stages = []struct {
	nodes  []node
	expect int
	bst    *BinaryTree
}{
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: 1},
			{id: "2", left: "4", right: "5", value: 2},
			{id: "3", left: "6", right: "7", value: 3},
			{id: "4", left: "8", right: "9", value: 4},
			{id: "5", left: "", right: "", value: 5},
			{id: "6", left: "", right: "", value: 6},
			{id: "7", left: "", right: "", value: 7},
			{id: "8", left: "", right: "", value: 8},
			{id: "9", left: "", right: "", value: 9},
		},
		expect: 16,
	},
	{
		nodes:  []node{{id: "1", left: "", right: "", value: 1}},
		expect: 0,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "", value: 1},
			{id: "2", left: "", right: "", value: 2},
		},
		expect: 1,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: 1},
			{id: "2", left: "", right: "", value: 2},
			{id: "3", left: "", right: "", value: 3},
		},
		expect: 2,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: 1},
			{id: "2", left: "4", right: "", value: 2},
			{id: "3", left: "", right: "", value: 3},
			{id: "4", left: "", right: "", value: 4},
		},
		expect: 4,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "", value: 1},
			{id: "2", left: "3", right: "", value: 2},
			{id: "3", left: "4", right: "", value: 3},
			{id: "4", left: "5", right: "", value: 4},
			{id: "5", left: "6", right: "", value: 5},
			{id: "6", left: "", right: "7", value: 6},
			{id: "7", left: "", right: "", value: 7},
		},
		expect: 21,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "8", value: 1},
			{id: "2", left: "3", right: "", value: 2},
			{id: "3", left: "4", right: "", value: 3},
			{id: "4", left: "5", right: "", value: 4},
			{id: "5", left: "6", right: "", value: 5},
			{id: "6", left: "", right: "7", value: 6},
			{id: "7", left: "", right: "", value: 7},
			{id: "8", left: "", right: "9", value: 8},
			{id: "9", left: "", right: "10", value: 9},
			{id: "10", left: "", right: "11", value: 10},
			{id: "11", left: "", right: "12", value: 11},
			{id: "12", left: "13", right: "", value: 12},
			{id: "13", left: "", right: "", value: 13},
		},
		expect: 42,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: 1},
			{id: "2", left: "4", right: "5", value: 2},
			{id: "3", left: "6", right: "7", value: 3},
			{id: "4", left: "8", right: "9", value: 4},
			{id: "5", left: "", right: "", value: 5},
			{id: "6", left: "10", right: "", value: 6},
			{id: "7", left: "", right: "", value: 7},
			{id: "8", left: "", right: "", value: 8},
			{id: "9", left: "", right: "", value: 9},
			{id: "10", left: "", right: "11", value: 10},
			{id: "11", left: "12", right: "13", value: 11},
			{id: "12", left: "14", right: "", value: 12},
			{id: "13", left: "15", right: "16", value: 13},
			{id: "14", left: "", right: "", value: 14},
			{id: "15", left: "", right: "", value: 15},
			{id: "16", left: "", right: "", value: 16},
		},
		expect: 51,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "", value: 1},
			{id: "2", left: "3", right: "", value: 2},
			{id: "3", left: "4", right: "", value: 3},
			{id: "4", left: "5", right: "", value: 4},
			{id: "5", left: "6", right: "", value: 5},
			{id: "6", left: "7", right: "", value: 6},
			{id: "7", left: "8", right: "", value: 7},
			{id: "8", left: "9", right: "", value: 8},
			{id: "9", left: "", right: "", value: 9},
		},
		expect: 36,
	},
}

func init() {
	for i := 0; i < len(stages); i++ {
		stages[i].bst = nodes2bst(stages[i].nodes)
	}
}

func TestBinaryTree(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := NodeDepths(stg.bst)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkBinaryTree(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		NodeDepths(stg.bst)
	}
}

func nodes2bst(nodes []node) *BinaryTree {
	reg := make(map[string]*node, len(nodes))
	for i := 0; i < len(nodes); i++ {
		reg[nodes[i].id] = &nodes[i]
	}
	node_ := &nodes[0]
	root := &BinaryTree{Value: node_.value}
	nodes2bst_(root, node_, nodes, reg)
	return root
}

func nodes2bst_(root *BinaryTree, node_ *node, nodes []node, reg map[string]*node) {
	if len(node_.left) > 0 {
		lnode := reg[node_.left]
		root.Left = &BinaryTree{Value: lnode.value}
		nodes2bst_(root.Left, lnode, nodes, reg)
	}
	if len(node_.right) > 0 {
		rnode := reg[node_.right]
		root.Right = &BinaryTree{Value: rnode.value}
		nodes2bst_(root.Right, rnode, nodes, reg)
	}
}

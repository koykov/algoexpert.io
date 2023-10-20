package binary_tree_diameter

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
			{id: "1", left: "3", right: "2", value: 1},
			{id: "3", left: "7", right: "4", value: 3},
			{id: "7", left: "8", right: "", value: 7},
			{id: "8", left: "9", right: "", value: 8},
			{id: "9", left: "", right: "", value: 9},
			{id: "4", left: "", right: "5", value: 4},
			{id: "5", left: "", right: "6", value: 5},
			{id: "6", left: "", right: "", value: 6},
			{id: "2", left: "", right: "", value: 2},
		},
		expect: 6,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: 1},
			{id: "3", left: "6", right: "7", value: 3},
			{id: "7", left: "", right: "", value: 7},
			{id: "6", left: "", right: "", value: 6},
			{id: "2", left: "4", right: "5", value: 2},
			{id: "5", left: "", right: "", value: 5},
			{id: "4", left: "", right: "", value: 4},
		},
		expect: 4,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: 1},
			{id: "3", left: "", right: "", value: 3},
			{id: "2", left: "", right: "", value: 2},
		},
		expect: 2,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "-1", value: 1},
			{id: "-1", left: "", right: "", value: -1},
			{id: "2", left: "", right: "", value: 2},
		},
		expect: 2,
	},
	{
		nodes: []node{
			{id: "1", left: "-5", right: "3", value: 1},
			{id: "3", left: "", right: "", value: 3},
			{id: "-5", left: "6", right: "", value: -5},
			{id: "6", left: "", right: "", value: 6},
		},
		expect: 3,
	},
	{
		nodes: []node{
			{id: "1", left: "3", right: "9", value: 1},
			{id: "3", left: "", right: "", value: 3},
			{id: "9", left: "14", right: "10", value: 9},
			{id: "10", left: "", right: "11", value: 10},
			{id: "11", left: "", right: "12", value: 11},
			{id: "12", left: "", right: "17", value: 12},
			{id: "17", left: "", right: "", value: 17},
			{id: "14", left: "", right: "19", value: 14},
			{id: "19", left: "25", right: "", value: 19},
			{id: "25", left: "", right: "", value: 25},
		},
		expect: 7,
	},
	{
		nodes: []node{
			{id: "1", left: "3", right: "5", value: 1},
			{id: "3", left: "", right: "", value: 3},
			{id: "5", left: "", right: "", value: 5},
		},
		expect: 2,
	},
	{
		nodes: []node{
			{id: "1", left: "5", right: "", value: 1},
			{id: "5", left: "7", right: "9", value: 5},
			{id: "9", left: "", right: "12", value: 9},
			{id: "12", left: "", right: "", value: 12},
			{id: "7", left: "8", right: "", value: 7},
			{id: "8", left: "", right: "", value: 8},
		},
		expect: 4,
	},
	{
		nodes: []node{
			{id: "1", left: "", right: "", value: 1},
		},
		expect: 0,
	},
	{
		nodes: []node{
			{id: "4", left: "2", right: "", value: 4},
			{id: "2", left: "", right: "", value: 2},
		},
		expect: 1,
	},
	{
		nodes: []node{
			{id: "4", left: "2", right: "", value: 4},
			{id: "2", left: "1", right: "", value: 2},
			{id: "1", left: "", right: "", value: 1},
		},
		expect: 2,
	},
	{
		nodes: []node{
			{id: "4", left: "2", right: "", value: 4},
			{id: "2", left: "1", right: "", value: 2},
			{id: "1", left: "", right: "3", value: 1},
			{id: "3", left: "19", right: "", value: 3},
			{id: "19", left: "", right: "", value: 19},
		},
		expect: 4,
	},
	{
		nodes: []node{
			{id: "6", left: "", right: "1", value: 6},
			{id: "1", left: "", right: "", value: 1},
		},
		expect: 1,
	},
	{
		nodes: []node{
			{id: "3", left: "", right: "10", value: 3},
			{id: "10", left: "1", right: "", value: 10},
			{id: "1", left: "", right: "", value: 1},
		},
		expect: 2,
	},
	{
		nodes: []node{
			{id: "2", left: "1", right: "", value: 2},
			{id: "1", left: "3", right: "", value: 1},
			{id: "3", left: "", right: "5", value: 3},
			{id: "5", left: "", right: "10", value: 5},
			{id: "10", left: "", right: "", value: 10},
		},
		expect: 4,
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
			r := BinaryTreeDiameter(stg.bst)
			if r != stg.expect {
				println(r)
				t.FailNow()
			}
		})
	}
}

func BenchmarkBinaryTree(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		BinaryTreeDiameter(stg.bst)
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

package evaluate_expression_tree

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
			{id: "1", left: "2", right: "3", value: -1},
			{id: "2", left: "", right: "", value: 2},
			{id: "3", left: "", right: "", value: 3},
		},
		expect: 5,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: -2},
			{id: "2", left: "", right: "", value: 2},
			{id: "3", left: "", right: "", value: 3},
		},
		expect: -1,
	},
	{
		nodes: []node{
			{id: "1", left: "3", right: "2", value: -2},
			{id: "2", left: "", right: "", value: 2},
			{id: "3", left: "", right: "", value: 3},
		},
		expect: 1,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: -3},
			{id: "2", left: "", right: "", value: 2},
			{id: "3", left: "", right: "", value: 3},
		},
		expect: 0,
	},
	{
		nodes: []node{
			{id: "1", left: "3", right: "2", value: -3},
			{id: "2", left: "", right: "", value: 2},
			{id: "3", left: "", right: "", value: 3},
		},
		expect: 1,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: -4},
			{id: "2", left: "", right: "", value: 2},
			{id: "3", left: "", right: "", value: 3},
		},
		expect: 6,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: -1},
			{id: "2", left: "", right: "", value: 2},
			{id: "3", left: "4", right: "5", value: -2},
			{id: "4", left: "", right: "", value: 4},
			{id: "5", left: "", right: "", value: 5},
		},
		expect: 1,
	},
	{
		nodes: []node{
			{id: "1", left: "10", right: "3", value: -3},
			{id: "10", left: "", right: "", value: 10},
			{id: "3", left: "4", right: "6", value: -2},
			{id: "4", left: "", right: "", value: 4},
			{id: "6", left: "", right: "", value: 6},
		},
		expect: -5,
	},
	{
		nodes: []node{
			{id: "1", left: "9", right: "3", value: -3},
			{id: "9", left: "", right: "", value: 9},
			{id: "3", left: "4", right: "6", value: -2},
			{id: "4", left: "", right: "", value: 4},
			{id: "6", left: "", right: "", value: 6},
		},
		expect: -4,
	},
	{
		nodes: []node{
			{id: "1", left: "9", right: "3", value: -3},
			{id: "9", left: "", right: "", value: 9},
			{id: "3", left: "4", right: "6", value: -2},
			{id: "4", left: "", right: "", value: 4},
			{id: "6", left: "", right: "", value: 6},
		},
		expect: -4,
	},
	{
		nodes: []node{
			{id: "1", left: "9", right: "3", value: -3},
			{id: "9", left: "", right: "", value: 9},
			{id: "3", left: "6", right: "4", value: -2},
			{id: "4", left: "", right: "", value: 4},
			{id: "6", left: "", right: "", value: 6},
		},
		expect: 4,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: -2},
			{id: "2", left: "4", right: "5", value: -1},
			{id: "3", left: "6", right: "7", value: -3},
			{id: "4", left: "", right: "", value: 7},
			{id: "5", left: "", right: "", value: 10},
			{id: "6", left: "", right: "", value: 12},
			{id: "7", left: "8", right: "9", value: -4},
			{id: "8", left: "", right: "", value: 1},
			{id: "9", left: "", right: "", value: 4},
		},
		expect: 14,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: -1},
			{id: "2", left: "4", right: "5", value: -2},
			{id: "3", left: "6", right: "7", value: -4},
			{id: "4", left: "", right: "", value: 7},
			{id: "5", left: "", right: "", value: 10},
			{id: "6", left: "", right: "", value: 12},
			{id: "7", left: "8", right: "9", value: -3},
			{id: "8", left: "", right: "", value: 8},
			{id: "9", left: "", right: "", value: 4},
		},
		expect: 21,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: -1},
			{id: "2", left: "4", right: "5", value: -1},
			{id: "3", left: "6", right: "7", value: -1},
			{id: "4", left: "", right: "", value: 7},
			{id: "5", left: "", right: "", value: 10},
			{id: "6", left: "", right: "", value: 12},
			{id: "7", left: "8", right: "9", value: -1},
			{id: "8", left: "", right: "", value: 8},
			{id: "9", left: "", right: "", value: 4},
		},
		expect: 41,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: -2},
			{id: "2", left: "4", right: "5", value: -2},
			{id: "3", left: "6", right: "7", value: -2},
			{id: "4", left: "", right: "", value: 7},
			{id: "5", left: "", right: "", value: 10},
			{id: "6", left: "", right: "", value: 12},
			{id: "7", left: "8", right: "9", value: -2},
			{id: "8", left: "", right: "", value: 8},
			{id: "9", left: "", right: "", value: 4},
		},
		expect: -11,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: -3},
			{id: "2", left: "4", right: "5", value: -3},
			{id: "3", left: "6", right: "7", value: -3},
			{id: "4", left: "", right: "", value: 100},
			{id: "5", left: "", right: "", value: 10},
			{id: "6", left: "", right: "", value: 4},
			{id: "7", left: "8", right: "9", value: -3},
			{id: "8", left: "", right: "", value: 8},
			{id: "9", left: "", right: "", value: 4},
		},
		expect: 5,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: -4},
			{id: "2", left: "4", right: "5", value: -4},
			{id: "3", left: "6", right: "7", value: -4},
			{id: "4", left: "", right: "", value: 2},
			{id: "5", left: "", right: "", value: 1},
			{id: "6", left: "", right: "", value: 7},
			{id: "7", left: "8", right: "9", value: -4},
			{id: "8", left: "", right: "", value: 8},
			{id: "9", left: "", right: "", value: 2},
		},
		expect: 224,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "9", value: -4},
			{id: "2", left: "4", right: "3", value: -1},
			{id: "3", left: "", right: "", value: 8},
			{id: "4", left: "5", right: "6", value: -1},
			{id: "5", left: "", right: "", value: 7},
			{id: "6", left: "7", right: "8", value: -2},
			{id: "7", left: "", right: "", value: 22},
			{id: "8", left: "", right: "", value: 5},
			{id: "9", left: "10", right: "11", value: -3},
			{id: "10", left: "", right: "", value: 100},
			{id: "11", left: "12", right: "13", value: -2},
			{id: "12", left: "", right: "", value: 42},
			{id: "13", left: "14", right: "15", value: -3},
			{id: "14", left: "16", right: "17", value: -4},
			{id: "15", left: "", right: "", value: 2},
			{id: "16", left: "", right: "", value: 3},
			{id: "17", left: "", right: "", value: 9},
		},
		expect: 96,
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
			r := EvaluateExpressionTree(stg.bst)
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
		EvaluateExpressionTree(stg.bst)
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

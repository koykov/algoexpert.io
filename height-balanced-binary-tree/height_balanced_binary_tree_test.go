package height_balanced_binary_tree

import (
	"strconv"
	"testing"
)

type node struct {
	id     string
	left   string
	right  string
	parent string
	value  int
}

var stages = []struct {
	nodes  []node
	expect bool
	tree   *BinaryTree
}{
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: 1},
			{id: "2", left: "4", right: "5", value: 2},
			{id: "3", left: "", right: "6", value: 3},
			{id: "4", left: "", right: "", value: 4},
			{id: "5", left: "7", right: "8", value: 5},
			{id: "6", left: "", right: "", value: 6},
			{id: "7", left: "", right: "", value: 7},
			{id: "8", left: "", right: "", value: 8},
		},
		expect: true,
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: 1},
			{id: "2", left: "4", right: "5", value: 2},
			{id: "3", left: "", right: "6", value: 3},
			{id: "4", left: "", right: "", value: 4},
			{id: "5", left: "7", right: "8", value: 5},
			{id: "6", left: "9", right: "10", value: 6},
			{id: "9", left: "", right: "", value: 9},
			{id: "10", left: "", right: "", value: 10},
			{id: "7", left: "", right: "", value: 7},
			{id: "8", left: "", right: "", value: 8},
		},
		expect: false,
	},
}

func init() {
	for i := 0; i < len(stages); i++ {
		stages[i].tree = nodes2bst(stages[i].nodes)
	}
}

func TestBinaryTree(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := HeightBalancedBinaryTree(stg.tree)
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
		HeightBalancedBinaryTree(stg.tree)
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

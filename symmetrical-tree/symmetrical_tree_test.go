package symmetrical_tree

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
	expect bool
	tree   *BinaryTree
}{
	{
		nodes: []node{
			{id: "1", left: "2", right: "2-2", value: 1},
			{id: "2", left: "3", right: "3-2", value: 2},
			{id: "2-2", left: "3-3", right: "3-4", value: 2},
			{id: "3", left: "", right: "", value: 3},
			{id: "3-2", left: "", right: "", value: 4},
			{id: "3-3", left: "", right: "", value: 4},
			{id: "3-4", left: "", right: "", value: 3},
		},
		expect: true,
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
			r := SymmetricalTree(stg.tree)
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
		SymmetricalTree(stg.tree)
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

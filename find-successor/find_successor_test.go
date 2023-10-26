package find_successor

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
	node_  int
	expect int
	tree   *BinaryTree
	node__ *BinaryTree
}{
	{
		nodes: []node{
			{id: "1", left: "2", parent: "", right: "3", value: 1},
			{id: "2", left: "4", parent: "1", right: "5", value: 2},
			{id: "3", left: "", parent: "1", right: "", value: 3},
			{id: "4", left: "6", parent: "2", right: "", value: 4},
			{id: "5", left: "", parent: "2", right: "", value: 5},
			{id: "6", left: "", parent: "4", right: "", value: 6},
		},
		node_:  5,
		expect: 1,
	},
	{
		nodes: []node{
			{id: "1", left: "2", parent: "", right: "3", value: 1},
			{id: "2", left: "4", parent: "1", right: "5", value: 2},
			{id: "3", left: "", parent: "1", right: "", value: 3},
			{id: "4", left: "", parent: "2", right: "", value: 4},
			{id: "5", left: "6", parent: "2", right: "7", value: 5},
			{id: "6", left: "", parent: "5", right: "", value: 6},
			{id: "7", left: "8", parent: "5", right: "", value: 7},
			{id: "8", left: "", parent: "7", right: "", value: 8},
		},
		node_:  5,
		expect: 8,
	},
	{
		nodes: []node{
			{id: "1", left: "2", parent: "", right: "", value: 1},
			{id: "2", left: "3", parent: "1", right: "", value: 2},
			{id: "3", left: "4", parent: "2", right: "", value: 3},
			{id: "4", left: "5", parent: "3", right: "", value: 4},
			{id: "5", left: "", parent: "4", right: "", value: 5},
		},
		node_:  3,
		expect: 2,
	},
	{
		nodes: []node{
			{id: "1", left: "2", parent: "", right: "5", value: 1},
			{id: "2", left: "", parent: "1", right: "3", value: 2},
			{id: "3", left: "", parent: "2", right: "4", value: 3},
			{id: "4", left: "", parent: "3", right: "", value: 4},
			{id: "5", left: "", parent: "1", right: "6", value: 5},
			{id: "6", left: "7", parent: "5", right: "8", value: 6},
			{id: "7", left: "", parent: "6", right: "", value: 7},
			{id: "8", left: "", parent: "6", right: "", value: 8},
		},
		node_:  1,
		expect: 5,
	},
}

func init() {
	for i := 0; i < len(stages); i++ {
		stages[i].tree, stages[i].node__ = nodes2bst(stages[i].nodes, stages[i].node_)
	}
}

func TestBinaryTree(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := FindSuccessor(stg.tree, stg.node__)
			if r.Value != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkBinaryTree(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		FindSuccessor(stg.tree, stg.node__)
	}
}

func nodes2bst(nodes []node, node1 int) (*BinaryTree, *BinaryTree) {
	reg := make(map[string]*node, len(nodes))
	for i := 0; i < len(nodes); i++ {
		reg[nodes[i].id] = &nodes[i]
	}
	node_ := &nodes[0]
	root := &BinaryTree{Value: node_.value}
	nodes2bst_(nil, root, node_, nodes, reg)
	return root, &BinaryTree{Value: reg[strconv.Itoa(node1)].value}
}

func nodes2bst_(parent, root *BinaryTree, node_ *node, nodes []node, reg map[string]*node) {
	if len(node_.left) > 0 {
		lnode := reg[node_.left]
		root.Left = &BinaryTree{Value: lnode.value}
		nodes2bst_(root, root.Left, lnode, nodes, reg)
	}
	root.Parent = parent
	if len(node_.right) > 0 {
		rnode := reg[node_.right]
		root.Right = &BinaryTree{Value: rnode.value}
		nodes2bst_(root, root.Right, rnode, nodes, reg)
	}
}

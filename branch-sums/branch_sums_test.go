package branch_sums

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
	expect []int
	bst    *BinaryTree
}{
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: 1},
			{id: "2", left: "4", right: "5", value: 2},
			{id: "3", left: "6", right: "7", value: 3},
			{id: "4", left: "8", right: "9", value: 4},
			{id: "5", left: "10", right: "", value: 5},
			{id: "6", left: "", right: "", value: 6},
			{id: "7", left: "", right: "", value: 7},
			{id: "8", left: "", right: "", value: 8},
			{id: "9", left: "", right: "", value: 9},
			{id: "10", left: "", right: "", value: 10},
		},
		expect: []int{15, 16, 18, 10, 11},
	},
	{
		nodes:  []node{{id: "1", left: "", right: "", value: 1}},
		expect: []int{1},
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "", value: 1},
			{id: "2", left: "", right: "", value: 2},
		},
		expect: []int{3},
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: 1},
			{id: "2", left: "", right: "", value: 2},
			{id: "3", left: "", right: "", value: 3},
		},
		expect: []int{3, 4},
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: 1},
			{id: "2", left: "4", right: "5", value: 2},
			{id: "3", left: "", right: "", value: 3},
			{id: "4", left: "", right: "", value: 4},
			{id: "5", left: "", right: "", value: 5},
		},
		expect: []int{7, 8, 4},
	},
	{
		nodes: []node{
			{id: "1", left: "2", right: "3", value: 1},
			{id: "2", left: "4", right: "5", value: 2},
			{id: "3", left: "6", right: "7", value: 3},
			{id: "4", left: "8", right: "9", value: 4},
			{id: "5", left: "10", right: "1-2", value: 5},
			{id: "6", left: "1-3", right: "1-4", value: 6},
			{id: "7", left: "", right: "", value: 7},
			{id: "8", left: "", right: "", value: 8},
			{id: "9", left: "", right: "", value: 9},
			{id: "10", left: "", right: "", value: 10},
			{id: "1-2", left: "", right: "", value: 1},
			{id: "1-3", left: "", right: "", value: 1},
			{id: "1-4", left: "", right: "", value: 1},
		},
		expect: []int{15, 16, 18, 9, 11, 11, 11},
	},
	{
		nodes: []node{
			{id: "0", left: "1", right: "", value: 0},
			{id: "1", left: "10", right: "", value: 1},
			{id: "10", left: "100", right: "", value: 10},
			{id: "100", left: "", right: "", value: 100},
		},
		expect: []int{111},
	},
	{
		nodes: []node{
			{id: "0", left: "", right: "1", value: 0},
			{id: "1", left: "", right: "10", value: 1},
			{id: "10", left: "", right: "100", value: 10},
			{id: "100", left: "", right: "", value: 100},
		},
		expect: []int{111},
	},
	{
		nodes: []node{
			{id: "0", left: "9", right: "1", value: 0},
			{id: "9", left: "", right: "", value: 9},
			{id: "1", left: "15", right: "10", value: 1},
			{id: "15", left: "", right: "", value: 15},
			{id: "10", left: "100", right: "200", value: 10},
			{id: "100", left: "", right: "", value: 100},
			{id: "200", left: "", right: "", value: 200},
		},
		expect: []int{9, 16, 111, 211},
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
			r := BranchSums(stg.bst)
			if !assert(r, stg.expect) {
				t.FailNow()
			}
		})
	}
}

func BenchmarkBinaryTree(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		BranchSums(stg.bst)
	}
}

func assert(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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

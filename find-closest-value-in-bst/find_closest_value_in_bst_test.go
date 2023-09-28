package find_closest_value_in_bst

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

var n0, n1 = []node{
	{id: "10", left: "5", right: "15", value: 10},
	{id: "15", left: "13", right: "22", value: 15},
	{id: "22", left: "", right: "", value: 22},
	{id: "13", left: "", right: "14", value: 13},
	{id: "14", left: "", right: "", value: 14},
	{id: "5", left: "2", right: "5-2", value: 5},
	{id: "5-2", left: "", right: "", value: 5},
	{id: "2", left: "1", right: "", value: 2},
	{id: "1", left: "", right: "", value: 1},
}, []node{
	{id: "100", left: "5", right: "502", value: 100},
	{id: "502", left: "204", right: "55000", value: 502},
	{id: "55000", left: "1001", right: "", value: 55000},
	{id: "1001", left: "", right: "4500", value: 1001},
	{id: "4500", left: "", right: "", value: 4500},
	{id: "204", left: "203", right: "205", value: 204},
	{id: "205", left: "", right: "207", value: 205},
	{id: "207", left: "206", right: "208", value: 207},
	{id: "208", left: "", right: "", value: 208},
	{id: "206", left: "", right: "", value: 206},
	{id: "203", left: "", right: "", value: 203},
	{id: "5", left: "2", right: "15", value: 5},
	{id: "15", left: "5-2", right: "22", value: 15},
	{id: "22", left: "", right: "57", value: 22},
	{id: "57", left: "", right: "60", value: 57},
	{id: "60", left: "", right: "", value: 60},
	{id: "5-2", left: "", right: "", value: 5},
	{id: "2", left: "1", right: "3", value: 2},
	{id: "3", left: "", right: "", value: 3},
	{id: "1", left: "-51", right: "1-2", value: 1},
	{id: "1-2", left: "", right: "1-3", value: 1},
	{id: "1-3", left: "", right: "1-4", value: 1},
	{id: "1-4", left: "", right: "1-5", value: 1},
	{id: "1-5", left: "", right: "", value: 1},
	{id: "-51", left: "-403", right: "", value: -51},
	{id: "-403", left: "", right: "", value: -403},
}
var stages = []struct {
	nodes  []node
	target int
	expect int
	bst    *BST
}{
	{
		nodes:  n0,
		target: 12,
		expect: 13,
	},
	{
		nodes:  n1,
		target: 100,
		expect: 100,
	},
	{
		nodes:  n1,
		target: 208,
		expect: 208,
	},
	{
		nodes:  n1,
		target: 4500,
		expect: 4500,
	},
	{
		nodes:  n1,
		target: 4501,
		expect: 4500,
	},
	{
		nodes:  n1,
		target: -70,
		expect: -51,
	},
	{
		nodes:  n1,
		target: 2000,
		expect: 1001,
	},
	{
		nodes:  n1,
		target: 6,
		expect: 5,
	},
	{
		nodes:  n1,
		target: 30000,
		expect: 55000,
	},
	{
		nodes:  n1,
		target: -1,
		expect: 1,
	},
	{
		nodes:  n1,
		target: 29751,
		expect: 55000,
	},
	{
		nodes:  n1,
		target: 29749,
		expect: 4500,
	},
}

func init() {
	for i := 0; i < len(stages); i++ {
		stages[i].bst = nodes2bst(stages[i].nodes)
	}
}

func TestBST(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := stg.bst.FindClosestValue(stg.target)
			if r != stg.expect {
				t.FailNow()
			}
		})
	}
}

func BenchmarkBTS(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		stg.bst.FindClosestValue(stg.target)
	}
}

func nodes2bst(nodes []node) *BST {
	reg := make(map[string]*node, len(nodes))
	for i := 0; i < len(nodes); i++ {
		reg[nodes[i].id] = &nodes[i]
	}
	node_ := &nodes[0]
	root := &BST{Value: node_.value}
	nodes2bst_(root, node_, nodes, reg)
	return root
}

func nodes2bst_(root *BST, node_ *node, nodes []node, reg map[string]*node) {
	if len(node_.left) > 0 {
		lnode := reg[node_.left]
		root.Left = &BST{Value: lnode.value}
		nodes2bst_(root.Left, lnode, nodes, reg)
	}
	if len(node_.right) > 0 {
		rnode := reg[node_.right]
		root.Right = &BST{Value: rnode.value}
		nodes2bst_(root.Right, rnode, nodes, reg)
	}
}

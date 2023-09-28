package depth_first_search

import (
	"strconv"
	"testing"
)

type node struct {
	id       string
	value    string
	children []string
}

var stages = []struct {
	nodes  []node
	expect []string
	root   *Node
}{
	{
		nodes: []node{
			{children: []string{"B", "C", "D"}, id: "A", value: "A"},
			{children: []string{"E", "F"}, id: "B", value: "B"},
			{children: []string{}, id: "C", value: "C"},
			{children: []string{"G", "H"}, id: "D", value: "D"},
			{children: []string{}, id: "E", value: "E"},
			{children: []string{"I", "J"}, id: "F", value: "F"},
			{children: []string{"K"}, id: "G", value: "G"},
			{children: []string{}, id: "H", value: "H"},
			{children: []string{}, id: "I", value: "I"},
			{children: []string{}, id: "J", value: "J"},
			{children: []string{}, id: "K", value: "K"},
		},
		expect: []string{"A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"},
	},
	{
		nodes: []node{
			{children: []string{"B", "C"}, id: "A", value: "A"},
			{children: []string{"D"}, id: "B", value: "B"},
			{children: []string{}, id: "C", value: "C"},
			{children: []string{}, id: "D", value: "D"},
		},
		expect: []string{"A", "B", "D", "C"},
	},
	{
		nodes: []node{
			{children: []string{"B", "C", "D", "E"}, id: "A", value: "A"},
			{children: []string{}, id: "B", value: "B"},
			{children: []string{"F"}, id: "C", value: "C"},
			{children: []string{}, id: "D", value: "D"},
			{children: []string{}, id: "E", value: "E"},
			{children: []string{}, id: "F", value: "F"},
		},
		expect: []string{"A", "B", "C", "F", "D", "E"},
	},
	{
		nodes: []node{
			{children: []string{"B"}, id: "A", value: "A"},
			{children: []string{"C"}, id: "B", value: "B"},
			{children: []string{"D", "E"}, id: "C", value: "C"},
			{children: []string{"F"}, id: "D", value: "D"},
			{children: []string{}, id: "E", value: "E"},
			{children: []string{}, id: "F", value: "F"},
		},
		expect: []string{"A", "B", "C", "D", "F", "E"},
	},
	{
		nodes: []node{
			{children: []string{"B", "C", "D", "E", "F"}, id: "A", value: "A"},
			{children: []string{"G", "H", "I"}, id: "B", value: "B"},
			{children: []string{"J"}, id: "C", value: "C"},
			{children: []string{"K", "L"}, id: "D", value: "D"},
			{children: []string{}, id: "E", value: "E"},
			{children: []string{"M", "N"}, id: "F", value: "F"},
			{children: []string{}, id: "G", value: "G"},
			{children: []string{"O", "P", "Q", "R"}, id: "H", value: "H"},
			{children: []string{}, id: "I", value: "I"},
			{children: []string{}, id: "J", value: "J"},
			{children: []string{"S"}, id: "K", value: "K"},
			{children: []string{}, id: "L", value: "L"},
			{children: []string{}, id: "M", value: "M"},
			{children: []string{}, id: "N", value: "N"},
			{children: []string{}, id: "O", value: "O"},
			{children: []string{"T", "U"}, id: "P", value: "P"},
			{children: []string{}, id: "Q", value: "Q"},
			{children: []string{"V"}, id: "R", value: "R"},
			{children: []string{}, id: "S", value: "S"},
			{children: []string{}, id: "T", value: "T"},
			{children: []string{}, id: "U", value: "U"},
			{children: []string{"W", "X", "Y"}, id: "V", value: "V"},
			{children: []string{}, id: "W", value: "W"},
			{children: []string{"Z"}, id: "X", value: "X"},
			{children: []string{}, id: "Y", value: "Y"},
			{children: []string{}, id: "Z", value: "Z"},
		},
		expect: []string{"A", "B", "G", "H", "O", "P", "T", "U", "Q", "R", "V", "W", "X", "Z", "Y", "I", "C", "J", "D", "K", "S", "L", "E", "F", "M", "N"},
	},
}

func init() {
	for i := 0; i < len(stages); i++ {
		stages[i].root = nodes2node(stages[i].nodes)
	}
}

func TestBinaryTree(t *testing.T) {
	eq := func(a, b []string) bool {
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
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var buf []string
			buf = stg.root.DepthFirstSearch(buf[:0])
			if !eq(buf, stg.expect) {
				t.FailNow()
			}
		})
	}
}

func BenchmarkBinaryTree(b *testing.B) {
	b.ReportAllocs()
	var buf []string
	for i := 0; i < b.N; i++ {
		stg := &stages[i%len(stages)]
		buf = stg.root.DepthFirstSearch(buf[:0])
	}
}

func nodes2node(nodes []node) *Node {
	reg := make(map[string]*node, len(nodes))
	for i := 0; i < len(nodes); i++ {
		reg[nodes[i].id] = &nodes[i]
	}
	node_ := &nodes[0]
	root := &Node{Name: node_.value}
	nodes2node_(root, node_, nodes, reg)
	return root
}

func nodes2node_(root *Node, node_ *node, nodes []node, reg map[string]*node) {
	if len(node_.children) > 0 {
		for i := 0; i < len(node_.children); i++ {
			c := reg[node_.children[i]]
			cn := &Node{Name: c.id}
			nodes2node_(cn, c, nodes, reg)
			root.Children = append(root.Children, cn)
		}
	}
}

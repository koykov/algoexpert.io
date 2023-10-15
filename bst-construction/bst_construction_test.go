package bst_construction

type node struct {
	id    string
	left  string
	right string
	value int
}

var stages = []struct {
	nodes  []node
	expect []int
	bst    *BST
}{}

// func TestBST(t *testing.T) {
// 	for i := 0; i < len(stages); i++ {
// 		stg := &stages[i]
// 		t.Run(strconv.Itoa(i), func(t *testing.T) {
// 			r := BranchSums(stg.bst)
// 			if !assert(r, stg.expect) {
// 				t.FailNow()
// 			}
// 		})
// 	}
// }
//
// func BenchmarkBST(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		stg := &stages[i%len(stages)]
// 		BranchSums(stg.bst)
// 	}
// }

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

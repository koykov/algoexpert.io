package find_nodes_distance_k

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func FindNodesDistanceK(tree *BinaryTree, target int, k int) (dst []int) {
	p := make(map[int]*BinaryTree)
	v := make(map[int]struct{})
	tnode := makep(p, tree, nil, target)
	if tnode == nil {
		return
	}
	dst = []int{}
	dst = walk(dst, p, v, tnode, k)
	return
}

func makep(dst map[int]*BinaryTree, node *BinaryTree, p *BinaryTree, target int) (tnode *BinaryTree) {
	if node == nil {
		return
	}
	if p != nil {
		dst[node.Value] = p
	}
	if node.Value == target {
		tnode = node
	}
	if tnl := makep(dst, node.Left, node, target); tnl != nil {
		tnode = tnl
	}
	if tnr := makep(dst, node.Right, node, target); tnr != nil {
		tnode = tnr
	}
	return
}

func walk(dst []int, p map[int]*BinaryTree, v map[int]struct{}, node *BinaryTree, k int) []int {
	if node == nil {
		return dst
	}
	if _, ok := v[node.Value]; ok {
		return dst
	}
	if k == 0 {
		dst = append(dst, node.Value)
		return dst
	}
	v[node.Value] = struct{}{}
	dst = walk(dst, p, v, node.Left, k-1)
	dst = walk(dst, p, v, node.Right, k-1)
	dst = walk(dst, p, v, p[node.Value], k-1)
	return dst
}

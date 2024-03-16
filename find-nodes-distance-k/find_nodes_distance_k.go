package find_nodes_distance_k

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func FindNodesDistanceK(tree *BinaryTree, target int, k int) []int {
	p := make(map[int]*BinaryTree)
	makep(p, tree, nil)
	return nil
}

func makep(dst map[int]*BinaryTree, tree *BinaryTree, p *BinaryTree) {
	if tree == nil {
		return
	}
	if p != nil {
		dst[tree.Value] = p
	}
	makep(dst, tree.Left, tree)
	makep(dst, tree.Right, tree)
}

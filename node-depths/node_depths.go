package node_depths

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	return dive(root, 0)
}

func dive(root *BinaryTree, depth int) (r int) {
	r += depth
	if root.Left != nil {
		r += dive(root.Left, depth+1)
	}
	if root.Right != nil {
		r += dive(root.Right, depth+1)
	}
	return
}

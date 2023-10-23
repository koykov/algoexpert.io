package find_successor

type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

func FindSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	var edge bool
	var r *BinaryTree
	r = walk(tree, node.Value, &edge, r)
	return r
}

func walk(tree *BinaryTree, value int, edge *bool, r *BinaryTree) *BinaryTree {
	if *edge {
		if r == nil {
			r = tree
		}
		return r
	}
	if tree.Left != nil {
		r = walk(tree.Left, value, edge, r)
	}
	if tree.Value == value {
		*edge = true
		if tree.Right != nil {
			r = tree.Right
		}
		return r
	}
	if *edge && r == nil {
		r = tree
	}
	if tree.Right != nil {
		r = walk(tree.Right, value, edge, r)
	}
	return r
}

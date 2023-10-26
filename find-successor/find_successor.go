package find_successor

type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

func FindSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	var ctl int
	var r *BinaryTree
	r = walk(tree, node.Value, &ctl, r)
	return r
}

func walk(tree *BinaryTree, value int, ctl *int, r *BinaryTree) *BinaryTree {
	if tree.Left != nil {
		r = walk(tree.Left, value, ctl, r)
	}
	if *ctl == -1 && r == nil {
		r = tree
		return r
	}
	if *ctl == 1 && r != nil {
		return r
	}
	if *ctl == 1 && tree.Left == nil {
		r = tree
		return r
	}
	if tree.Value == value {
		if tree.Right != nil {
			*ctl = 1
		} else {
			*ctl = -1
			return r
		}
	}
	if tree.Right != nil {
		r = walk(tree.Right, value, ctl, r)
	}
	return r
}

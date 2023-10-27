package merge_binary_trees

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func MergeBinaryTrees(t0 *BinaryTree, t1 *BinaryTree) *BinaryTree {
	node := BinaryTree{}
	var l0, l1, r0, r1 *BinaryTree
	if t0 != nil {
		node.Value += t0.Value
		l0, r0 = t0.Left, t0.Right
	}
	if t1 != nil {
		node.Value += t1.Value
		l1, r1 = t1.Left, t1.Right
	}
	if l0 != nil || l1 != nil {
		node.Left = MergeBinaryTrees(l0, l1)
	}
	if r0 != nil || r1 != nil {
		node.Right = MergeBinaryTrees(r0, r1)
	}
	return &node
}

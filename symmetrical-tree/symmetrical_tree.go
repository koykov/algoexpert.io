package symmetrical_tree

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func SymmetricalTree(tree *BinaryTree) bool {
	reverse(tree.Right)
	return equal(tree.Left, tree.Right)
}

func reverse(tree *BinaryTree) {
	if tree == nil {
		return
	}
	reverse(tree.Left)
	reverse(tree.Right)
	tree.Left, tree.Right = tree.Right, tree.Left
}

func equal(t0, t1 *BinaryTree) bool {
	if t0 == nil && t1 == nil {
		return true
	}
	if (t0 != nil && t1 == nil) || (t0 == nil && t1 != nil) {
		return false
	}
	leq := equal(t0.Left, t1.Left)
	req := equal(t0.Right, t1.Right)
	return t0.Value == t1.Value && leq && req
}

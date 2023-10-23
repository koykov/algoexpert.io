package binary_tree_diameter

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func BinaryTreeDiameter(tree *BinaryTree) (d int) {
	walk(tree, &d)
	return
}

func walk(root *BinaryTree, d *int) int {
	if root == nil {
		return 0
	}
	lh, rh := walk(root.Left, d), walk(root.Right, d)
	*d = max(*d, lh+rh)
	return max(lh, rh) + 1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

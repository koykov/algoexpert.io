package height_balanced_binary_tree

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func HeightBalancedBinaryTree(tree *BinaryTree) bool {
	l, lok := walk(tree.Left)
	r, rok := walk(tree.Right)
	return abs(l-r) <= 1 && lok && rok
}

func walk(tree *BinaryTree) (int, bool) {
	if tree == nil {
		return 0, true
	}
	if tree.Left == nil && tree.Right == nil {
		return 1, true
	}
	l, lok := walk(tree.Left)
	r, rok := walk(tree.Right)
	return max(l, r) + 1, abs(l-r) <= 1 && lok && rok
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

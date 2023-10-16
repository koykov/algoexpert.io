package validate_bst

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) ValidateBst() bool {
	return check(tree, -math.MaxInt, math.MaxInt)
}

func check(root *BST, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Value < min || root.Value >= max {
		return false
	}
	return check(root.Left, min, root.Value) && check(root.Right, root.Value, max)
}

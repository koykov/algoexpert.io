package split_binary_tree

import "math"

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func SplitBinaryTree(tree *BinaryTree) int {
	t := walk(tree, -math.MaxInt, nil)
	if t%2 != 0 {
		return 0
	}
	var ok bool
	walk(tree, t/2, &ok)
	if ok {
		return t / 2
	}
	return 0
}

func walk(tree *BinaryTree, target int, ok *bool) (r int) {
	r += tree.Value
	if tree.Left != nil {
		r += walk(tree.Left, target, ok)
	}
	if tree.Right != nil {
		r += walk(tree.Right, target, ok)
	}
	if target != -math.MaxInt && r == target {
		*ok = true
		return
	}
	return
}

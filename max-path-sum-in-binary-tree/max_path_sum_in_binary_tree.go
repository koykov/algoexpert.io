package max_path_sum_in_binary_tree

import "math"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func MaxPathSum(tree *BinaryTree) int {
	buf := make([]int, 0)
	var l, r int
	buf, l = sumOf(buf, tree.Left)
	buf, r = sumOf(buf, tree.Right)
	buf = append(buf, l+tree.Value)
	buf = append(buf, r+tree.Value)
	buf = append(buf, l+r+tree.Value)
	m := -math.MaxInt
	for i := 0; i < len(buf); i++ {
		if buf[i] > m {
			m = buf[i]
		}
	}
	return m
}

func sumOf(buf []int, tree *BinaryTree) ([]int, int) {
	if tree == nil {
		return buf, 0
	}
	var l, r int
	buf, l = sumOf(buf, tree.Left)
	buf, r = sumOf(buf, tree.Right)
	buf = append(buf, l+r+tree.Value)
	if l > r {
		buf = append(buf, l+tree.Value)
		return buf, l + tree.Value
	} else {
		buf = append(buf, r+tree.Value)
		return buf, r + tree.Value
	}
}

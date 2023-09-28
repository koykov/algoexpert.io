package find_closest_value_in_bst

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) (res int) {
	if tree.Value-target == 0 {
		return tree.Value
	}
	l, r := math.MaxInt/2, math.MaxInt/2
	if tree.Left != nil {
		l = tree.Left.FindClosestValue(target)
	}
	if tree.Right != nil {
		r = tree.Right.FindClosestValue(target)
	}
	ld, rd, cd := abs(l-target), abs(r-target), abs(tree.Value-target)
	mn := rd
	res = r
	if ld < rd {
		res, mn = l, ld
	}
	if cd < mn {
		res = tree.Value
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

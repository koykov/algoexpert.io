package find_kth_largest_value_in_bst

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func FindKthLargestValueInBst(tree *BST, k int) (r int) {
	walkReverse(tree, &k, &r)
	return
}

func walkReverse(root *BST, k, r *int) {
	if root.Right != nil {
		walkReverse(root.Right, k, r)
	}
	*k--
	if *k == 0 {
		*r = root.Value
	}
	if root.Left != nil {
		walkReverse(root.Left, k, r)
	}
}

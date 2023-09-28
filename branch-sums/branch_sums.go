package branch_sums

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	buf := make([]int, 0, 32)
	return sum(buf, root, 0)
}

func sum(dst []int, root *BinaryTree, acc int) []int {
	acc += root.Value
	if root.Left == nil && root.Right == nil {
		dst = append(dst, acc)
		return dst
	}
	if root.Left != nil {
		dst = sum(dst, root.Left, acc)
	}
	if root.Right != nil {
		dst = sum(dst, root.Right, acc)
	}
	return dst
}

package binary_tree_diameter

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func BinaryTreeDiameter(tree *BinaryTree) (r int) {
	var c int
	walk(tree, &c, &r)
	return r
}

func walk(tree *BinaryTree, c, r *int) {
	if tree.Left == nil && tree.Right == nil {
		if *r < *c {
			*r = *c
		}
		*c = 0
		return
	}
	if tree.Left != nil {
		walk(tree.Left, c, r)
	}
	if tree.Right != nil {
		walk(tree.Right, c, r)
	}
}

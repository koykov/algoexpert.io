package bst_traversal

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) InOrderTraverse(dst []int) []int {
	if tree.Left != nil {
		dst = tree.Left.InOrderTraverse(dst)
	}
	dst = append(dst, tree.Value)
	if tree.Right != nil {
		dst = tree.Right.InOrderTraverse(dst)
	}
	return dst
}

func (tree *BST) PreOrderTraverse(dst []int) []int {
	dst = append(dst, tree.Value)
	if tree.Left != nil {
		dst = tree.Left.PreOrderTraverse(dst)
	}
	if tree.Right != nil {
		dst = tree.Right.PreOrderTraverse(dst)
	}
	return dst
}

func (tree *BST) PostOrderTraverse(dst []int) []int {
	if tree.Left != nil {
		dst = tree.Left.PostOrderTraverse(dst)
	}
	if tree.Right != nil {
		dst = tree.Right.PostOrderTraverse(dst)
	}
	dst = append(dst, tree.Value)
	return dst
}

package bst_construction

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	return insert(tree, value)
}

func (tree *BST) Contains(value int) bool {
	if tree.Value == value {
		return true
	}
	if tree.Left != nil && tree.Left.Contains(value) {
		return true
	}
	if tree.Right != nil && tree.Right.Contains(value) {
		return true
	}
	return false
}

func (tree *BST) Remove(value int) *BST {
	root := remove(tree, value)
	if root == nil {
		return nil
	}
	*tree = *root
	return tree
}

func insert(x *BST, value int) *BST {
	if x == nil {
		x = &BST{Value: value}
	} else if value < x.Value {
		x.Left = insert(x.Left, value)
	} else if value >= x.Value {
		x.Right = insert(x.Right, value)
	}
	return x
}

func remove(x *BST, value int) *BST {
	if x == nil {
		return x
	}
	switch {
	case value < x.Value:
		x.Left = remove(x.Left, value)
	case value > x.Value:
		x.Right = remove(x.Right, value)
	case x.Left != nil && x.Right != nil:
		x.Value = minimum(x.Right).Value
		x.Right = remove(x.Right, x.Value)
	default:
		switch {
		case x.Left != nil:
			x = x.Left
		case x.Right != nil:
			x = x.Right
		default:
			x = nil
		}
	}
	return x
}

func minimum(root *BST) *BST {
	switch {
	case root.Left == nil && root.Right == nil:
		return root
	case root.Left == nil && root.Right != nil:
		return minimum(root.Right)
	case root.Left != nil:
		return minimum(root.Left)
	}
	return root
}

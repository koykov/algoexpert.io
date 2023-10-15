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
	return remove(tree, value)
}

func insert(x *BST, value int) *BST {
	if x == nil {
		x = &BST{Value: value}
	} else if value < x.Value {
		insert(x.Left, value)
	} else if value > x.Value {
		insert(x.Right, value)
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
		x.Value = x.Right.Value
		x.Right = remove(x.Right, x.Value)
	default:
		if x.Left != nil {
			x = x.Left
		} else if x.Right != nil {
			x = x.Right
		} else {
			x = nil
		}
	}
	return x
}

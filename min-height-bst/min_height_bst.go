package min_height_bst

import "strconv"

func MinHeightBST(a []int) *BST {
	var root *BST
	switch len(a) {
	case 0:
		return root
	case 1:
		root = &BST{Value: a[0]}
	case 2:
		root = &BST{Value: a[1]}
		root.Left = &BST{Value: a[0]}
	default:
		p := len(a) / 2
		root = &BST{Value: a[p]}
		root.Left = MinHeightBST(a[:p])
		root.Right = MinHeightBST(a[p+1:])
	}
	return root
}

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}

func (tree *BST) Marshal() string {
	out := "{value:" + strconv.Itoa(tree.Value)
	if tree.Left != nil {
		out += ",left:" + tree.Left.Marshal()
	}
	if tree.Right != nil {
		out += ",right:" + tree.Right.Marshal()
	}
	out += "}"
	return out
}

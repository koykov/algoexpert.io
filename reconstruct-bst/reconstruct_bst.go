package reconstruct_bst

import "strconv"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func ReconstructBst(a []int) *BST {
	if len(a) == 0 {
		return nil
	}
	root := &BST{Value: a[0]}
	ri := -1
	for i := 1; i < len(a); i++ {
		if a[i] >= root.Value {
			ri = i
			break
		}
	}
	switch {
	case ri == -1:
		root.Left = ReconstructBst(a[1:])
	default:
		root.Left = ReconstructBst(a[1:ri])
		root.Right = ReconstructBst(a[ri:])
	}
	return root
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

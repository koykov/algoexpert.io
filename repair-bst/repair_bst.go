package repair_bst

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func RepairBst(tree *BST) *BST {
	var n1, n2, np *BST
	n1, n2, np = walk(tree, n1, n2, np)
	n1.Value, n2.Value = n2.Value, n1.Value
	return tree
}

func walk(root, n1, n2, np *BST) (*BST, *BST, *BST) {
	if root == nil {
		return n1, n2, np
	}
	n1, n2, np = walk(root.Left, n1, n2, np)
	if np != nil && np.Value > root.Value {
		if n1 == nil {
			n1 = np
		}
		n2 = root
	}
	np = root
	n1, n2, np = walk(root.Right, n1, n2, np)
	return n1, n2, np
}

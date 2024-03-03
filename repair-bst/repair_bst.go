package repair_bst

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func RepairBst(tree *BST) *BST {
	_, lnode, _ := walk(tree.Value, tree, tree.Left, true)
	_, rnode, _ := walk(tree.Value, tree, tree.Right, false)

	lnode.Value, rnode.Value = rnode.Value, lnode.Value
	return tree
}

func walk(rv int, root, node *BST, less bool) (*BST, *BST, bool) {
	if node == nil {
		return root, node, false
	}
	if less {
		if node.Value > rv {
			return root, node, true
		}
	} else {
		if node.Value < rv {
			return root, node, true
		}
	}
	if rr, rn, ok := walk(rv, node, node.Left, true); ok {
		return rr, rn, ok
	}
	if rr, rn, ok := walk(rv, node, node.Right, false); ok {
		return rr, rn, ok
	}
	return root, node, false
}

package repair_bst

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func RepairBst(tree *BST) *BST {
	lnode, lok := walk(tree.Value, tree.Left, true)
	rnode, rok := walk(tree.Value, tree.Right, false)

	if lok && rok {
		lnode.Value, rnode.Value = rnode.Value, lnode.Value
	}
	return tree
}

func walk(rv int, node *BST, less bool) (*BST, bool) {
	if node == nil {
		return node, false
	}
	if less {
		if node.Value > rv {
			return node, true
		}
	} else {
		if node.Value < rv {
			return node, true
		}
	}
	lnode, lok := walk(rv, node.Left, true)
	rnode, rok := walk(rv, node.Right, false)
	if lok && rok {
		lnode.Value, rnode.Value = rnode.Value, lnode.Value
		return node, false
	}
	if lok {
		return lnode, true
	}
	if rok {
		return rnode, true
	}
	return node, false
}

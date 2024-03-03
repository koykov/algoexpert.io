package validate_three_nodes

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func ValidateThreeNodes(n1, n2, n3 *BST) bool {
	if walk(n1, n2) {
		return walk(n2, n3)
	}
	if walk(n3, n2) {
		return walk(n2, n1)
	}
	return false
}

func walk(root, ch *BST) bool {
	if root.Value == ch.Value {
		return true
	}
	var l, r bool
	if root.Left != nil {
		l = walk(root.Left, ch)
	}
	if root.Right != nil {
		r = walk(root.Right, ch)
	}
	return l || r
}

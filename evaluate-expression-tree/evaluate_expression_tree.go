package evaluate_expression_tree

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func EvaluateExpressionTree(tree *BinaryTree) int {
	if tree.Left == nil && tree.Right == nil {
		return tree.Value
	}
	var l, r int
	if tree.Left != nil {
		l = EvaluateExpressionTree(tree.Left)
	}
	if tree.Right != nil {
		r = EvaluateExpressionTree(tree.Right)
	}
	switch tree.Value {
	case -1:
		return l + r
	case -2:
		return l - r
	case -3:
		return l / r
	case -4:
		return l * r
	default:
		panic(tree.Value)
	}
}

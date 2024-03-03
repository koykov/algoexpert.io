package repair_bst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepairBst(t *testing.T) {
	tree := &BST{Value: 2}
	tree.Left = &BST{Value: 1}
	tree.Right = &BST{Value: 3}
	tree.Left.Left = &BST{Value: 4}
	tree.Right.Right = &BST{Value: 0}
	expected := []int{0, 1, 2, 3, 4}
	actual := inOrderTraverse(RepairBst(tree), nil)
	require.Equal(t, expected, actual)
}

func inOrderTraverse(tree *BST, array []int) []int {
	if tree == nil {
		return array
	}
	array = inOrderTraverse(tree.Left, array)
	array = append(array, tree.Value)
	array = inOrderTraverse(tree.Right, array)
	return array
}

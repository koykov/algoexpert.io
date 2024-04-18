package sum_bsts

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumBsts(t *testing.T) {
	root := &BinaryTree{Value: 8}
	root.Left = &BinaryTree{Value: 2}
	root.Left.Left = &BinaryTree{Value: 1}
	root.Left.Right = &BinaryTree{Value: 10}
	root.Right = &BinaryTree{Value: 9}
	root.Right.Right = &BinaryTree{Value: 5}
	expected := 13
	actual := SumBsts(root)
	require.Equal(t, expected, actual)
}

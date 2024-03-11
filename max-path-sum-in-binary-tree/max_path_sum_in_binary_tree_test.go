package max_path_sum_in_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxPathSum(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		test := NewBinaryTree(1).insertAll([]int{2, 3, 4, 5, 6, 7})
		require.Equal(t, MaxPathSum(test), 18)
	})
	t.Run("1", func(t *testing.T) {
		test := NewBinaryTree(1).insertAll([]int{2, -1})
		require.Equal(t, MaxPathSum(test), 3)
	})
}

func NewBinaryTree(value int) *BinaryTree {
	return &BinaryTree{Value: value}
}

func (tree *BinaryTree) insertAll(values []int) *BinaryTree {
	for _, value := range values {
		tree.insert(value)
	}
	return tree
}

func (tree *BinaryTree) insert(value int) {
	queue := []*BinaryTree{tree}
	var current *BinaryTree
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if current.Left == nil {
			current.Left = NewBinaryTree(value)
			break
		}
		queue = append(queue, current.Left)
		if current.Right == nil {
			current.Right = NewBinaryTree(value)
			break
		}
		queue = append(queue, current.Right)
	}
}

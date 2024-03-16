package find_nodes_distance_k

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindNodesDistanceK(t *testing.T) {
	root := &BinaryTree{Value: 1}
	root.Left = &BinaryTree{Value: 2}
	root.Right = &BinaryTree{Value: 3}
	root.Left.Left = &BinaryTree{Value: 4}
	root.Left.Right = &BinaryTree{Value: 5}
	root.Right.Right = &BinaryTree{Value: 6}
	root.Right.Right.Left = &BinaryTree{Value: 7}
	root.Right.Right.Right = &BinaryTree{Value: 8}
	target := 3
	k := 2
	expected := []int{2, 7, 8}
	actual := FindNodesDistanceK(root, target, k)
	sort.Ints(actual)
	require.Equal(t, expected, actual)
}

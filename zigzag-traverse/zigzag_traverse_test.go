package zigzag_traverse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZigzagTraverse(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		input := [][]int{
			{1, 3, 4, 10},
			{2, 5, 9, 11},
			{6, 8, 12, 15},
			{7, 13, 14, 16},
		}
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
		output := ZigzagTraverse(input)
		require.Equal(t, expected, output)
	})
	t.Run("1", func(t *testing.T) {
		input := [][]int{
			{1, 2, 3, 4, 5},
		}
		expected := []int{1, 2, 3, 4, 5}
		output := ZigzagTraverse(input)
		require.Equal(t, expected, output)
	})
}

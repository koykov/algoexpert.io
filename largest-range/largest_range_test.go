package largest_range

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLargestRange(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		expected := []int{0, 7}
		output := LargestRange([]int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6})
		require.Equal(t, expected, output)
	})
	t.Run("1", func(t *testing.T) {
		expected := []int{1, 1}
		output := LargestRange([]int{1})
		require.Equal(t, expected, output)
	})
}

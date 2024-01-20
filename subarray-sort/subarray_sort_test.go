package subarray_sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubarraySort(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		expected := []int{3, 9}
		output := SubarraySort([]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19})
		require.Equal(t, expected, output)
	})
	t.Run("1", func(t *testing.T) {
		expected := []int{4, 6}
		output := SubarraySort([]int{1, 2, 4, 7, 10, 11, 7, 12, 13, 14, 16, 18, 19})
		require.Equal(t, expected, output)
	})
}

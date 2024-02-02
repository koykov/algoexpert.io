package longest_subarray_with_sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestSubarrayWithSum(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 3, 3, 1, 2, 1}
		targetSum := 10
		expected := []int{4, 8}
		actual := LongestSubarrayWithSum(array, targetSum)
		require.Equal(t, expected, actual)
	})
	t.Run("1", func(t *testing.T) {
		array := []int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0}
		targetSum := 1
		expected := []int{0, 9}
		actual := LongestSubarrayWithSum(array, targetSum)
		require.Equal(t, expected, actual)
	})
	t.Run("2", func(t *testing.T) {
		array := []int{0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 5, 5, 5, 5, 1, 1, 0, 1, 0, 1, 0, 0, 0, 1}
		targetSum := 20
		expected := []int{11, 23}
		actual := LongestSubarrayWithSum(array, targetSum)
		require.Equal(t, expected, actual)
	})
	t.Run("3", func(t *testing.T) {
		array := []int{1, 4, 10, 15, 31, 7, 1, 40, 0, 20, 1, 1, 1, 1, 2, 1}
		targetSum := 0
		expected := []int{8, 8}
		actual := LongestSubarrayWithSum(array, targetSum)
		require.Equal(t, expected, actual)
	})
}

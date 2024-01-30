package longest_subarray_with_sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestSubarrayWithSum(t *testing.T) {
	array := []int{1, 2, 3, 4, 3, 3, 1, 2, 1}
	targetSum := 10
	expected := []int{4, 8}
	actual := LongestSubarrayWithSum(array, targetSum)
	require.Equal(t, expected, actual)
}

package max_sum_increasing_subsequence

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxSumIncreasingSubsequence(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		outputSum, outputSequence := MaxSumIncreasingSubsequence([]int{10, 70, 20, 30, 50, 11, 30})
		require.Equal(t, 110, outputSum)
		require.Equal(t, []int{10, 20, 30, 50}, outputSequence)
	})
	t.Run("1", func(t *testing.T) {
		outputSum, outputSequence := MaxSumIncreasingSubsequence([]int{-1})
		require.Equal(t, -1, outputSum)
		require.Equal(t, []int{-1}, outputSequence)
	})
	t.Run("2", func(t *testing.T) {
		outputSum, outputSequence := MaxSumIncreasingSubsequence([]int{-1, 1})
		require.Equal(t, 1, outputSum)
		require.Equal(t, []int{1}, outputSequence)
	})
	t.Run("3", func(t *testing.T) {
		outputSum, outputSequence := MaxSumIncreasingSubsequence([]int{10, 15, 4, 5, 11, 14, 31, 25, 31, 23, 25, 31, 50})
		require.Equal(t, 164, outputSum)
		require.Equal(t, []int{10, 11, 14, 23, 25, 31, 50}, outputSequence)
	})
}

package search_for_range

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchForRange(t *testing.T) {
	t.Run("", func(t *testing.T) {
		expected := []int{4, 9}
		output := SearchForRange([]int{0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 61, 71, 73}, 45)
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := []int{0, 0}
		output := SearchForRange([]int{5, 7, 7, 8, 8, 10}, 5)
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := []int{-1, -1}
		output := SearchForRange([]int{5, 7, 7, 8, 8, 10}, 9)
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := []int{1, 2}
		output := SearchForRange([]int{5, 7, 7, 8, 8, 10}, 7)
		require.Equal(t, expected, output)
	})
}

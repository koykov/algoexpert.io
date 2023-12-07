package permutations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPermutations(t *testing.T) {
	output := GetPermutations([]int{1, 2, 3})
	require.Contains(t, output, []int{1, 2, 3})
	require.Contains(t, output, []int{1, 3, 2})
	require.Contains(t, output, []int{2, 1, 3})
	require.Contains(t, output, []int{2, 3, 1})
	require.Contains(t, output, []int{3, 1, 2})
	require.Contains(t, output, []int{3, 2, 1})
	require.Len(t, output, 6)
}

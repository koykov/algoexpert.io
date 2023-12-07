package powerset

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPowerset(t *testing.T) {
	output := Powerset([]int{1, 2, 3})
	require.Contains(t, output, []int{})
	require.Contains(t, output, []int{1})
	require.Contains(t, output, []int{2})
	require.Contains(t, output, []int{1, 2})
	require.Contains(t, output, []int{3})
	require.Contains(t, output, []int{1, 3})
	require.Contains(t, output, []int{2, 3})
	require.Contains(t, output, []int{1, 2, 3})
	require.Len(t, output, 8)
}

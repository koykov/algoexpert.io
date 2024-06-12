package shifted_binary_search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShiftedBinarySearch(t *testing.T) {
	expected := 8
	output := ShiftedBinarySearch([]int{45, 61, 71, 72, 73, 0, 1, 21, 33, 37}, 33)
	require.Equal(t, expected, output)
}

package maximum_sum_submatrix

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaximumSumSubmatrix(t *testing.T) {
	matrix := [][]int{
		{5, 3, -1, 5},
		{-7, 3, 7, 4},
		{12, 8, 0, 0},
		{1, -8, -8, 2},
	}
	size := 2
	expected := 18
	actual := MaximumSumSubmatrix(matrix, size)
	require.Equal(t, expected, actual)
}

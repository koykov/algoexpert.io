package sort_stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortStack(t *testing.T) {
	input := []int{-5, 2, -2, 4, 3, 1}
	expected := []int{-5, -2, 1, 2, 3, 4}
	actual := SortStack(input)
	require.Equal(t, expected, actual)
}

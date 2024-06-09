package quick_sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuickSort(t *testing.T) {
	expected := []int{2, 3, 5, 5, 6, 8, 9}
	output := QuickSort([]int{8, 5, 2, 9, 5, 6, 3})
	require.Equal(t, expected, output)
}

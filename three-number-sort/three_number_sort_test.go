package three_number_sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThreeNumberSort(t *testing.T) {
	array := []int{1, 0, 0, -1, -1, 0, 1, 1}
	order := []int{0, 1, -1}
	expected := []int{0, 0, 0, 1, 1, 1, -1, -1}
	actual := ThreeNumberSort(array, order)
	require.Equal(t, expected, actual)
}

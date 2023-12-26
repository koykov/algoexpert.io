package next_greater_element

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNextGreaterElement(t *testing.T) {
	input := []int{2, 5, -3, -4, 6, 7, 2}
	expected := []int{5, 6, 6, 6, 7, -1, 5}
	actual := NextGreaterElement(input)
	require.Equal(t, expected, actual)
}

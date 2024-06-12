package index_equals_value

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndexEqualsValue(t *testing.T) {
	input := []int{-5, -3, 0, 3, 4, 5, 9}
	expected := 3
	actual := IndexEqualsValue(input)
	require.Equal(t, expected, actual)
}

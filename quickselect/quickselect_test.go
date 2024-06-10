package quickselect

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuickselect(t *testing.T) {
	expected := 5
	output := Quickselect([]int{8, 5, 2, 9, 7, 6, 3}, 3)
	require.Equal(t, expected, output)
}

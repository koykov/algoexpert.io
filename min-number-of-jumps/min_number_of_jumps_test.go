package min_number_of_jumps

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinNumberOfJumps(t *testing.T) {
	expected := 4
	output := MinNumberOfJumps([]int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3})
	require.Equal(t, expected, output)
}

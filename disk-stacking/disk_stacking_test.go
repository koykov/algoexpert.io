package disk_stacking

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiskStacking(t *testing.T) {
	expected := [][]int{{2, 1, 2}, {3, 2, 3}, {4, 4, 5}}
	input := [][]int{{2, 1, 2}, {3, 2, 3}, {2, 2, 8}, {2, 3, 4}, {2, 2, 1}, {4, 4, 5}}
	output := DiskStacking(input)
	require.Equal(t, expected, output)
}

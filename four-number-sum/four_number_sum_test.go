package four_number_sum

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, expected, output [][]int) {
	t.Helper()
	require.Len(t, output, len(expected))
	for _, quad := range expected {
		sort.Ints(quad)
		ourquad := fmt.Sprintf("%v", quad)
		found := false
		for _, theirquad := range output {
			sort.Ints(theirquad)
			if fmt.Sprintf("%v", theirquad) == ourquad {
				found = true
				break
			}
		}
		require.True(t, found)
	}
}

func TestFourNumberSum(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		expected := [][]int{{7, 6, 4, -1}, {7, 6, 1, 2}}
		output := FourNumberSum([]int{7, 6, 4, -1, 1, 2}, 16)
		doTest(t, expected, output)
	})
	t.Run("1", func(t *testing.T) {
		expected := [][]int{{1, 2, 3, 4}}
		output := FourNumberSum([]int{1, 2, 3, 4, 5, 6, 7}, 10)
		doTest(t, expected, output)
	})
}

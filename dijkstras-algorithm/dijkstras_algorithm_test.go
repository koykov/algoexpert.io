package dijkstras_algorithm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDijkstrasAlgorithm(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		start := 0
		edges := [][][]int{
			{{1, 7}},
			{{2, 6}, {3, 20}, {4, 3}},
			{{3, 14}},
			{{4, 2}},
			{},
			{},
		}
		expected := []int{0, 7, 13, 27, 10, -1}
		actual := DijkstrasAlgorithm(start, edges)
		require.Equal(t, expected, actual)
	})
	t.Run("1", func(t *testing.T) {
		start := 0
		edges := [][][]int{
			{{1, 1}, {3, 1}},
			{{2, 1}},
			{{6, 1}},
			{{1, 3}, {2, 4}, {4, 2}, {5, 3}, {6, 5}},
			{{5, 1}},
			{{4, 1}},
			{{5, 2}},
			{{0, 7}},
		}
		expected := []int{0, 7, 13, 27, 10, -1}
		actual := DijkstrasAlgorithm(start, edges)
		require.Equal(t, expected, actual)
	})
}

package count_squares

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountSquares(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		input := [][]int{{1, 1}, {0, 0}, {-4, 2}, {-2, -1}, {0, 1}, {1, 0}, {-1, 4}}
		expected := 2
		actual := CountSquares(input)
		require.Equal(t, expected, actual)
	})
	t.Run("1", func(t *testing.T) {
		input := [][]int{{1, 1}, {0, 0}, {0, 1}, {1, 0}}
		expected := 1
		actual := CountSquares(input)
		require.Equal(t, expected, actual)
	})
}

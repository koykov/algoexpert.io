package knight_connection

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKnightConnection(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		knightA := []int{0, 0}
		knightB := []int{2, 1}
		expected := 1
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("1", func(t *testing.T) {
		knightA := []int{15, -12}
		knightB := []int{15, -12}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
}

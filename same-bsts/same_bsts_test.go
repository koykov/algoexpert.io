package same_bsts

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSameBsts(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		arrayOne := []int{10, 15, 8, 12, 94, 81, 5, 2, 11}
		arrayTwo := []int{10, 8, 5, 15, 2, 12, 11, 94, 81}
		require.Equal(t, SameBsts(arrayOne, arrayTwo), true)
	})
	t.Run("1", func(t *testing.T) {
		arrayOne := []int{5, 2, -1, 100, 45, 12, 8, -1, 8, 10, 15, 8, 12, 94, 81, 2, -34}
		arrayTwo := []int{5, 8, 10, 15, 2, 8, 12, 45, 100, 2, 12, 94, 81, -1, -1, -34, 8}
		require.Equal(t, SameBsts(arrayOne, arrayTwo), false)
	})
	t.Run("2", func(t *testing.T) {
		arrayOne := []int{10, 15, 8, 12, 94, 81, 5, 2, -1, 100, 45, 12, 9, -1, 8, 2, -34}
		arrayTwo := []int{10, 8, 5, 15, 2, 12, 94, 81, -1, -1, -34, 8, 2, 9, 12, 45, 100}
		require.Equal(t, SameBsts(arrayOne, arrayTwo), false)
	})
	t.Run("3", func(t *testing.T) {
		arrayOne := []int{10, 15, 8, 12, 94, 81, 5, 2, 10}
		arrayTwo := []int{10, 8, 5, 15, 2, 10, 12, 94, 81}
		require.Equal(t, SameBsts(arrayOne, arrayTwo), false)
	})
}

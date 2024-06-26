package water_area

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWaterArea(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		expected := 48
		output := WaterArea([]int{0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3})
		require.Equal(t, expected, output)
	})
	t.Run("1", func(t *testing.T) {
		expected := 1
		output := WaterArea([]int{0, 1, 0, 1, 0})
		require.Equal(t, expected, output)
	})
}

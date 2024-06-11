package longest_most_frequent_prefix

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestMostFrequentPrefix(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		strings := []string{"algoexpert", "algorithm", "frontendexpert", "mlexpert"}
		expected := "algo"
		actual := LongestMostFrequentPrefix(strings)
		require.Equal(t, expected, actual)
	})
	t.Run("1", func(t *testing.T) {
		strings := []string{"hello", "world", "fossil", "worldly", "food", "forrest", "helium", "algoexpert", "algorithm"}
		expected := "fo"
		actual := LongestMostFrequentPrefix(strings)
		require.Equal(t, expected, actual)
	})
}

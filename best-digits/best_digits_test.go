package best_digits

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBestDigits(t *testing.T) {
	t.Run("462839-2", func(t *testing.T) {
		number := "462839"
		numDigits := 2
		expected := "6839"
		actual := BestDigits(number, numDigits)
		require.Equal(t, expected, actual)
	})
	t.Run("22-1", func(t *testing.T) {
		number := "22"
		numDigits := 1
		expected := "2"
		actual := BestDigits(number, numDigits)
		require.Equal(t, expected, actual)
	})
	t.Run("321-1", func(t *testing.T) {
		number := "321"
		numDigits := 1
		expected := "32"
		actual := BestDigits(number, numDigits)
		require.Equal(t, expected, actual)
	})
}

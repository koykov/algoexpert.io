package reversePolishNotation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReversePolishNotation(t *testing.T) {
	t.Run("35", func(t *testing.T) {
		input := []string{"3", "2", "+", "7", "*"}
		expected := 35
		actual := ReversePolishNotation(input)
		require.Equal(t, expected, actual)
	})
	t.Run("10", func(t *testing.T) {
		input := []string{"10"}
		expected := 10
		actual := ReversePolishNotation(input)
		require.Equal(t, expected, actual)
	})
	t.Run("42", func(t *testing.T) {
		input := []string{"4", "-7", "2", "6", "+", "10", "-", "/", "*", "2", "+", "3", "*"}
		expected := 42
		actual := ReversePolishNotation(input)
		require.Equal(t, expected, actual)
	})
}

package longest_palindromic_substring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	t.Run("abaxyzzyxf", func(t *testing.T) {
		expected := "xyzzyx"
		output := LongestPalindromicSubstring("abaxyzzyxf")
		require.Equal(t, expected, output)
	})
	t.Run("it's highnoon", func(t *testing.T) {
		expected := "noon"
		output := LongestPalindromicSubstring("it's highnoon")
		require.Equal(t, expected, output)
	})
	t.Run("abcdefgfedcbazzzzzzzzzzzzzzzzzzzz", func(t *testing.T) {
		expected := "zzzzzzzzzzzzzzzzzzzz"
		output := LongestPalindromicSubstring("abcdefgfedcbazzzzzzzzzzzzzzzzzzzz")
		require.Equal(t, expected, output)
	})
	t.Run("z234a5abbba54a32z", func(t *testing.T) {
		expected := "5abbba5"
		output := LongestPalindromicSubstring("z234a5abbba54a32z")
		require.Equal(t, expected, output)
	})
}

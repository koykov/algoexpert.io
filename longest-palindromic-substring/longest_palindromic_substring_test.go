package longest_palindromic_substring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	expected := "xyzzyx"
	output := LongestPalindromicSubstring("abaxyzzyxf")
	require.Equal(t, expected, output)
}

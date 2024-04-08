package longest_common_subsequence

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestCommonSubsequence(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		expected := []byte("XYZW")
		output := LongestCommonSubsequence("ZXVVYZW", "XKYKZPW")
		require.Equal(t, expected, output)
	})
	t.Run("1", func(t *testing.T) {
		expected := []byte("AE")
		output := LongestCommonSubsequence("ABCDEFG", "APPLES")
		require.Equal(t, expected, output)
	})
}

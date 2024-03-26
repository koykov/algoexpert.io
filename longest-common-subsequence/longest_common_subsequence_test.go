package longest_common_subsequence

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestCommonSubsequence(t *testing.T) {
	expected := []byte("XYZW")
	output := LongestCommonSubsequence("ZXVVYZW", "XKYKZPW")
	require.Equal(t, expected, output)
}

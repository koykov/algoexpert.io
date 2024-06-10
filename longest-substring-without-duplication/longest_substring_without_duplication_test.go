package longest_substring_without_duplication

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestSubstringWithoutDuplication(t *testing.T) {
	expected := "mentisac"
	output := LongestSubstringWithoutDuplication("clementisacap")
	require.Equal(t, expected, output)
}

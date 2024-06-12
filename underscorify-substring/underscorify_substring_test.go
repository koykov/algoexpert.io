package underscorify_substring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnderscorifySubstring(t *testing.T) {
	t.Run("", func(t *testing.T) {
		expected := "_test_this is a _testtest_ to see if _testestest_ it works"
		output := UnderscorifySubstring("testthis is a testtest to see if testestest it works", "test")
		require.Equal(t, expected, output)
	})
	t.Run("", func(t *testing.T) {
		expected := "_tttttttttttttt_b_ttttt_ctatawta_ttttt_astvb"
		output := UnderscorifySubstring("ttttttttttttttbtttttctatawtatttttastvb", "ttt")
		require.Equal(t, expected, output)
	})
}

package reverse_words_in_string

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseWordsInString(t *testing.T) {
	t.Run("AlgoExpert is the best!", func(t *testing.T) {
		input := "AlgoExpert is the best!"
		expected := "best! the is AlgoExpert"
		actual := ReverseWordsInString(input)
		require.Equal(t, expected, actual)
	})
	t.Run("..H,, hello 678", func(t *testing.T) {
		input := "..H,, hello 678"
		expected := "678 hello ..H,,"
		actual := ReverseWordsInString(input)
		require.Equal(t, expected, actual)
	})
	t.Run("this-is-one-word", func(t *testing.T) {
		input := "this-is-one-word"
		expected := "this-is-one-word"
		actual := ReverseWordsInString(input)
		require.Equal(t, expected, actual)
	})
	t.Run("this      string     has a     lot of   whitespace", func(t *testing.T) {
		input := "this      string     has a     lot of   whitespace"
		expected := "whitespace   of lot     a has     string      this"
		actual := ReverseWordsInString(input)
		require.Equal(t, expected, actual)
	})
}

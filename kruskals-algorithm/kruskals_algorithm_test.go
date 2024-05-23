package kruskals_algorithm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKruskalsAlgorithm(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		input := [][][]int{{{1, 1}}, {{0, 1}}}
		expected := [][][]int{{{1, 1}}, {{0, 1}}}
		actual := KruskalsAlgorithm(input)
		require.Equal(t, expected, actual)
	})
	// t.Run("1", func(t *testing.T) {
	// 	input := [][][]int{{{1, 1}}, {{0, 1}}}
	// 	expected := [][][]int{{{1, 1}}, {{0, 1}}}
	// 	actual := KruskalsAlgorithm(input)
	// 	require.Equal(t, expected, actual)
	// })
}

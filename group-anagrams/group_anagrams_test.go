package group_anagrams

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroupAnagrams(t *testing.T) {
	words := []string{"yo", "act", "flop", "tac", "foo", "cat", "oy", "olfp"}
	expected := [][]string{{"yo", "oy"}, {"flop", "olfp"}, {"act", "tac", "cat"}, {"foo"}}
	output := GroupAnagrams(words)
	compare(t, expected, output)
}

func compare(t *testing.T, expected, output [][]string) {
	t.Helper()
	for _, group := range output {
		sort.Strings(group)
	}

	for _, group := range expected {
		sort.Strings(group)
	}
	require.ElementsMatch(t, expected, output)
}

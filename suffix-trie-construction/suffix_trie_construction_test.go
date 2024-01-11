package suffix_trie_construction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func (trie SuffixTrie) Equals(other SuffixTrie) bool {
	if len(trie) != len(other) {
		return false
	}
	for key, child := range trie {
		otherchild, found := other[key]
		if !found {
			return false
		} else if child != nil && !child.Equals(otherchild) {
			return false
		}
	}
	return true
}

func TrieFromString(str string) SuffixTrie {
	trie := SuffixTrie{}
	trie.PopulateSuffixTrieFrom(str)
	return trie
}

func TestNewSuffixTrie(t *testing.T) {
	t.Run("babc", func(t *testing.T) {
		trie := TrieFromString("babc")
		expected := SuffixTrie{
			'c': {'*': nil},
			'b': {
				'c': {'*': nil},
				'a': {'b': {'c': {'*': nil}}},
			},
			'a': {'b': {'c': {'*': nil}}},
		}
		require.True(t, trie.Equals(expected))
		require.True(t, trie.Contains("abc"))
	})
	t.Run("1234556789", func(t *testing.T) {
		trie := TrieFromString("1234556789")
		expected := SuffixTrie{
			'c': {'*': nil},
			'b': {
				'c': {'*': nil},
				'a': {'b': {'c': {'*': nil}}},
			},
			'a': {'b': {'c': {'*': nil}}},
		}
		require.True(t, trie.Equals(expected))
		require.True(t, trie.Contains("abc"))
	})
}

package suffix_trie_construction

type SuffixTrie map[byte]SuffixTrie

func NewSuffixTrie() SuffixTrie {
	trie := SuffixTrie{}
	return trie
}

func (trie SuffixTrie) PopulateSuffixTrieFrom(s string) {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := trie[c]; !ok {
			trie[c] = NewSuffixTrie()
			trie[c]['*'] = nil
		}
		trie[c].PopulateSuffixTrieFrom(s[i+1:])
	}
}

func (trie SuffixTrie) Contains(s string) bool {
	return false
}

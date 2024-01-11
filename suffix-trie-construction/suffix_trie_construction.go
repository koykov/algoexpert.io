package suffix_trie_construction

type SuffixTrie map[byte]SuffixTrie

func NewSuffixTrie() SuffixTrie {
	trie := SuffixTrie{}
	return trie
}

func (trie SuffixTrie) PopulateSuffixTrieFrom(s string) {
	for i := 0; i < len(s); i++ {
		trie.Insert(s[i:])
	}
}

func (trie SuffixTrie) Insert(s string) {
	node := trie
	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := node[c]; !ok {
			node[c] = NewSuffixTrie()
		}
		node = node[c]
	}
	node['*'] = nil
}

func (trie SuffixTrie) Contains(s string) bool {
	node := trie
	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := node[c]; !ok {
			return false
		}
		node = node[c]
	}
	_, ok := node['*']
	return ok
}

func (trie SuffixTrie) String() string {
	buf := make([]byte, 0, 1024)
	buf = trie.write(buf)
	return string(buf)
}

func (trie SuffixTrie) write(buf []byte) []byte {
	buf = append(buf, '{')
	for k, v := range trie {
		buf = append(buf, k)
		buf = append(buf, ':')
		if v == nil {
			buf = append(buf, "nil"...)
		} else {
			buf = v.write(buf)
		}
	}
	buf = append(buf, '}')
	return buf
}

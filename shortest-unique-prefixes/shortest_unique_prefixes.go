package shortest_unique_prefixes

func ShortestUniquePrefixes(ss []string) []string {
	t := trie{root: &node{childs: make(map[string]*node)}}
	for i := 0; i < len(ss); i++ {
		t.insert(ss[i])
	}
	var r []string
	for i := 0; i < len(ss); i++ {
		r = append(r, t.sup(ss[i]))
	}
	return r
}

type node struct {
	childs map[string]*node
	visit  int
}

type trie struct {
	root *node
}

func (t *trie) insert(s string) {
	n := t.root
	for _, c := range s {
		cs := string(c)
		ch, ok := n.childs[cs]
		if !ok {
			ch = &node{childs: make(map[string]*node)}
			n.childs[cs] = ch
		}
		n = ch
		n.visit++
	}
}

func (t *trie) sup(s string) string {
	n := t.root
	for i, c := range s {
		cs := string(c)
		ch, ok := n.childs[cs]
		if !ok {
			return ""
		}
		n = ch
		if n.visit == 1 {
			return s[0 : i+1]
		}
	}
	return s
}

package shortest_unique_prefixes

func ShortestUniquePrefixes(ss []string) []string {
	r := make(map[byte][]string)
	for i := 0; i < len(ss); i++ {
		r[ss[i][0]] = append(r[ss[i][0]], ss[i])
	}
	var o []string
	for _, ss1 := range r {
		ch := lcp(ss1)
		o = append(o, ch...)
	}
	return o
}

func lcp(ss []string) []string {
	if len(ss) == 1 {
		return []string{ss[0][:1]}
	}
	var r []string
	s := ss[0]
	for i := 1; i < len(ss); i++ {
		for j := 1; j < mn(len(s), len(ss[i])); j++ {
			if j == len(ss[i]) {
				s = ss[i]
				break
			}
			if s[j] != ss[i][j] {
				s = s[0:j]
				break
			}
		}
	}
	for i := 0; i < len(ss); i++ {
		if len(ss[i]) > len(s) {
			r = append(r, ss[i][:len(s)+1])
		}
		if len(ss[i]) == len(s) {
			r = append(r, s)
		}
	}
	return r
}

func mn(a, b int) int {
	if a < b {
		return a
	}
	return b
}

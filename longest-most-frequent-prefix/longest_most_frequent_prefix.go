package longest_most_frequent_prefix

func LongestMostFrequentPrefix(ss []string) string {
	if len(ss) == 1 {
		return ss[0]
	}
	r := make(map[byte][]string)
	for i := 0; i < len(ss); i++ {
		r[ss[i][0]] = append(r[ss[i][0]], ss[i])
	}
	var ln string
	var c int
	for _, ss1 := range r {
		ln1 := lcp(ss1)
		if len(ss1) > c || (len(ss1) == c && len(ln1) > len(ln)) {
			ln, c = ln1, len(ss1)
		}
	}
	return ln
}

func lcp(ss []string) string {
	r := ss[0]
	for i := 1; i < len(ss); i++ {
		for j := 1; j < mn(len(r), len(ss[i])); j++ {
			if j == len(ss[i]) {
				r = ss[i]
				break
			}
			if r[j] != ss[i][j] {
				r = r[0:j]
				break
			}
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

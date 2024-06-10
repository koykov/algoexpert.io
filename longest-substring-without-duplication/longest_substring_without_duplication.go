package longest_substring_without_duplication

func LongestSubstringWithoutDuplication(s string) string {
	n := len(s)
	idx := make(map[int]int)
	tidx := make(map[byte]int)
	for i := 0; i < n; i++ {
		if _, ok := tidx[s[i]]; ok {
			idx[tidx[s[i]]] = i
		}
		tidx[s[i]] = i
	}
	for i := 0; i < n; i++ {
		if _, ok := idx[i]; !ok {
			idx[i] = n
		}
	}
	l, r, e, ln := 0, 1, n, 0
	for i := 0; i < n; i++ {
		e = n
		for j := i; j < e; j++ {
			e = mn(e, idx[j])
		}
		if ln <= (e - i - 1) {
			ln = e - i - 1
			l = i
			r = e
		}
	}
	return s[l:r]
}

func mn(a, b int) int {
	if a < b {
		return a
	}
	return b
}

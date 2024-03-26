package longest_common_subsequence

func LongestCommonSubsequence(s1 string, s2 string) []byte {
	mn, mx := s1, s2
	if len(mn) > len(mx) {
		mn, mx = mx, mn
	}
	r := make(map[byte]struct{}, len(mn))
	for i := 0; i < len(mn); i++ {
		r[mn[i]] = struct{}{}
	}
	buf := make([]byte, 0, len(mn))
	for i := 0; i < len(mx); i++ {
		if _, ok := r[mx[i]]; ok {
			buf = append(buf, mx[i])
		}
	}
	return buf
}

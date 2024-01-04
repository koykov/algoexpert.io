package longest_palindromic_substring

func LongestPalindromicSubstring(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	var l, r int
	f := func(s string, l, r int) (int, int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		return l + 1, r
	}
	for i := 0; i < n; i++ {
		ol, or := f(s, i, i)
		el, er := f(s, i, i+1)
		if d := or - ol; d > r-l {
			l, r = ol, or
		}
		if d := er - el; d > r-l {
			l, r = el, er
		}
	}
	return s[l:r]
}

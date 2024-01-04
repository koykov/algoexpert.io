package longest_palindromic_substring

func LongestPalindromicSubstring(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	var (
		mp string
		ml int
	)
	for i := 0; i < n; i++ {
		l, r := i, i
		if r < n-1 && s[i] == s[i+1] {
			r++
		}
		for l > 0 && r < n-1 {
			l--
			r++
			if s[l] != s[r] {
				break
			}
		}
		ds := s[l : r+1]
		if len(ds) > 1 && ds[0] != ds[len(ds)-1] {
			ds = ds[1 : len(ds)-1]
		}
		if len(ds) > ml {
			ml = len(ds)
			mp = s[l : r+1]
		}
	}
	if len(mp) > 1 && mp[0] != mp[len(mp)-1] {
		mp = mp[1 : len(mp)-1]
	}
	return mp
}

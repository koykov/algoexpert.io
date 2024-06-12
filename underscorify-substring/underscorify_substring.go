package underscorify_substring

import "strings"

func UnderscorifySubstring(s string, ss string) string {
	n := len(ss)
	buf := make([]byte, 0, len(s)*2)
	for {
		p := strings.Index(s, ss)
		if p == -1 {
			buf = append(buf, s...)
			break
		}
		if p > 0 {
			buf = append(buf, s[:p]...)
		}
		buf = append(buf, '_')
		for {
			pp := strings.Index(s[p+1:], ss)
			if pp == -1 || pp+1 > n {
				buf = append(buf, ss...)
				break
			}
			var d int
			if pp > 0 && pp+1 < n {
				d = n - pp - 1
			}
			buf = append(buf, ss[:n-d]...)
			s = s[p+pp+1:]
			p = 0
		}
		buf = append(buf, '_')
		s = s[p+len(ss):]
	}
	return string(buf)
}

package underscorify_substring

import "strings"

func UnderscorifySubstring(s string, ss string) string {
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
		if len(buf) > 0 && buf[len(buf)-1] == '_' {
			buf = buf[:len(buf)-1]
		} else {
			buf = append(buf, '_')
		}
		buf = append(buf, ss...)
		buf = append(buf, '_')
		pp := strings.Index(s[p+1:], ss)
		if pp > 0 && pp-p+1 < len(ss) {
			buf = buf[:len(buf)-pp-1]
			s = s[p+pp+1:]
		} else {
			s = s[p+len(ss):]
		}
	}
	return string(buf)
}

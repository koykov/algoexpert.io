package shorten_path

import "strings"

func ShortenPath(p string) string {
	if p == "/" || len(p) == 0 {
		return p
	}
	var stk []string
	if p[0] == '/' {
		stk = append(stk, "")
		p = p[1:]
	} else {
		for {
			t := p[:2]
			if t == ".." {
				stk = append(stk, t)
				p = p[3:]
				continue
			}
			break
		}
	}

	var t, pt string
	for {
		t, p = token(p)
		switch t {
		case ".", "":
			// CWD, do nothing
		case "..":
			if len(stk) > 0 && stk[len(stk)-1] != "" {
				stk = stk[:len(stk)-1]
			}
			// switch {
			// case len(stk) == 1 && stk[0] != "" && stk[0] != "..":
			// 	stk = stk[:len(stk)-1]
			// case len(stk) > 1:
			// 	stk = stk[:len(stk)-1]
			// default:
			// 	stk = append(stk, t)
			// }
		default:
			stk = append(stk, t)
		}
		pt = t
		_ = pt
		if len(p) == 0 {
			break
		}
	}
	return strings.Join(stk, "/")
}

func token(s string) (string, string) {
	n := len(s)
	switch {
	case n > 2 && s[:2] == "..":
		return s[:2], s[3:]
	case n > 1 && s[:1] == ".":
		return s[:1], s[2:]
	default:
		p := strings.IndexByte(s, '/')
		if p == -1 {
			return s, ""
		}
		return s[:p], s[p+1:]
	}
}

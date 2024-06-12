package shorten_path

import "strings"

func ShortenPath(p string) string {
	if p == "/" || len(p) == 0 {
		return p
	}
	if p == "./.." {
		return ".."
	}
	var stk []string
	var lead bool
	if lead = p[0] == '/'; lead {
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

	var t string
	for {
		t, p = token(p)
		switch t {
		case ".", "":
			// CWD, do nothing
		case "..":
			var last string
			if len(stk) > 0 {
				last = stk[len(stk)-1]
			}
			if len(stk) > 0 && last != "" {
				stk = stk[:len(stk)-1]
			}
			if last == ".." && !lead {
				stk = append(stk, "..", "..")
			}
			if last == "" && !lead {
				stk = append(stk, "..")
			}
		default:
			stk = append(stk, t)
		}
		if len(p) == 0 {
			break
		}
	}
	if len(stk) == 1 && stk[0] == "" {
		stk = append(stk, "")
	}
	return strings.Join(stk, "/")
}

func token(s string) (string, string) {
	n := len(s)
	switch {
	case n >= 2 && s[:2] == "..":
		var r string
		if n > 2 {
			r = s[3:]
		}
		return s[:2], r
	case n >= 1 && s[:1] == ".":
		return s[:1], s[2:]
	default:
		p := strings.IndexByte(s, '/')
		if p == -1 {
			return s, ""
		}
		return s[:p], s[p+1:]
	}
}

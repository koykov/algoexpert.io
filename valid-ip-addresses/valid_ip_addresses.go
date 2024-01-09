package valid_ip_addresses

import "strconv"

func ValidIPAddresses(s string) []string {
	r := make([]string, 0)
	if len(s) < 4 {
		return r
	}

	var buf []byte
	r, buf = try(r, buf[:0], s[:1], s[1:], 1)
	r, buf = try(r, buf[:0], s[:2], s[2:], 1)
	r, buf = try(r, buf[:0], s[:3], s[3:], 1)

	return r
}

func try(dst []string, buf []byte, o, s string, depth int) ([]string, []byte) {
	if depth == 4 {
		if len(s) == 0 && valid(o) {
			buf = append(buf, '.')
			buf = append(buf, o...)
			dst = append(dst, string(buf))
			return dst, buf
		}
		return dst, buf
	}
	if !valid(o) {
		return dst, buf
	}
	if depth > 1 {
		buf = append(buf, '.')
	}
	buf = append(buf, o...)
	off := len(buf)
	if len(s) > 0 {
		dst, buf = try(dst, buf[:off], s[:1], s[1:], depth+1)
	}
	if len(s) > 1 {
		dst, buf = try(dst, buf[:off], s[:2], s[2:], depth+1)
	}
	if len(s) > 2 {
		dst, buf = try(dst, buf[:off], s[:3], s[3:], depth+1)
	}
	return dst, buf
}

func valid(s string) bool {
	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if i == 0 && len(s) > 1 {
		return false
	}
	if i > 0 && s[0] == '0' {
		return false
	}
	return i >= 0 && i <= 255
}

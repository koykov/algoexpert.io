package run_length_encoding

import "strconv"

func RunLengthEncoding(str string) string {
	p := []byte(str)
	buf := make([]byte, 0, len(p)*2)
	x, c := p[0], 1
	for i := 1; i < len(p); i++ {
		if p[i] == x {
			c++
			continue
		}
		buf = out(buf, x, c)
		x, c = p[i], 1
	}
	buf = out(buf, x, c)
	return string(buf)
}

func out(dst []byte, x byte, c int) []byte {
	for c > 9 {
		dst = append(dst, '9')
		dst = append(dst, x)
		c = c - 9
	}
	if c > 0 {
		dst = strconv.AppendInt(dst, int64(c), 10)
		dst = append(dst, x)
	}
	return dst
}

package semordnilap

func Semordnilap(words []string) [][]string {
	reg := make(map[string]struct{}, len(words))
	for i := 0; i < len(words); i++ {
		reg[words[i]] = struct{}{}
	}
	r := make([][]string, 0, len(words)/2+1)
	for s := range reg {
		p := palindrome(s)
		if s == p {
			continue
		}
		if _, ok := reg[p]; ok {
			r = append(r, []string{s, p})
			delete(reg, s)
			delete(reg, p)
		}
	}
	return r
}

func palindrome(s string) string {
	buf := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		buf = append(buf, s[i])
	}
	return string(buf)
}

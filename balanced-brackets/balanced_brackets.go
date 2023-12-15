package balanced_brackets

func BalancedBrackets(s string) bool {
	var p byte
	f := func(e byte) bool {
		return p != e && p != ')' && p != ']' && p != '}'
	}

	var b, q, c int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			b++
			p = s[i]
		case '[':
			q++
			p = s[i]
		case '{':
			c++
			p = s[i]
		case ')':
			if i == 0 || f('(') {
				return false
			}
			b--
			p = s[i]
		case ']':
			if i == 0 || f('[') {
				return false
			}
			q--
			p = s[i]
		case '}':
			if i == 0 || f('{') {
				return false
			}
			c--
			p = s[i]
		}
	}
	return b == 0 && q == 0 && c == 0
}

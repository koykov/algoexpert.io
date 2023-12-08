package phone_number_mnemonics

var m = map[byte][]byte{
	'0': {'0'},
	'1': {'1'},
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func PhoneNumberMnemonics(pn string) (r []string) {
	if len(pn) == 0 {
		return []string{}
	}
	r = jonny(r, pn, "")
	return
}

func jonny(dst []string, pn, r string) []string {
	if len(pn) == 0 {
		dst = append(dst, r)
		return dst
	}
	if mn, ok := m[pn[0]]; ok {
		for i := 0; i < len(mn); i++ {
			r += string(mn[i])
			dst = jonny(dst, pn[1:], r)
			r = r[:len(r)-1]
		}
	}
	return dst
}

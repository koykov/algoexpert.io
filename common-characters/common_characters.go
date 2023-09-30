package common_characters

import "strings"

func CommonCharacters(ss []string) []string {
	reg := make(map[byte]struct{}, 100)
	for i := 0; i < len(ss[0]); i++ {
		reg[ss[0][i]] = struct{}{}
	}
	for c := range reg {
		for i := 1; i < len(ss); i++ {
			if strings.IndexByte(ss[i], c) == -1 {
				delete(reg, c)
			}
		}
	}
	r := make([]string, 0, len(reg))
	for c := range reg {
		r = append(r, string(c))
	}
	return r
}

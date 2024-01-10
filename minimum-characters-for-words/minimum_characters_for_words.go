package minimum_characters_for_words

func MinimumCharactersForWords(words []string) []string {
	r := make([]string, 0)
	idx, tidx := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(words); i++ {
		for k := range idx {
			delete(idx, k)
		}
		w := words[i]
		for j := 0; j < len(w); j++ {
			idx[w[j]]++
		}
		for k, v := range idx {
			if tidx[k] < v {
				tidx[k] = v
			}
		}
	}
	for b, c := range tidx {
		s := string(b)
		for i := 0; i < c; i++ {
			r = append(r, s)
		}
	}
	return r
}

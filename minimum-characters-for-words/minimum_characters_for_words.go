package minimum_characters_for_words

func MinimumCharactersForWords(words []string) []string {
	r := make([]string, 0)
	idx := make(map[byte]int)
	for i := 0; i < len(words); i++ {
		w := words[i]
		for j := 0; j < len(w); j++ {
			idx[w[j]]++
		}
	}
	for b, c := range idx {
		s := string(b)
		for i := 0; i < c; i++ {
			r = append(r, s)
		}
	}
	return r
}

package reverse_words_in_string

func ReverseWordsInString(s string) string {
	buf := make([]byte, 0, len(s))
	prev := len(s)
	i := prev - 1
	for i = len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && prev != i {
			buf = append(buf, s[i+1:prev]...)
			buf = append(buf, ' ')
			prev = i
		}
	}
	if i != prev {
		buf = append(buf, s[i+1:prev]...)
	}
	return string(buf)
}

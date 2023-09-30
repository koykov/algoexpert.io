package generate_document

func GenerateDocument(characters string, document string) bool {
	reg := make(map[byte]int, len(characters))
	for i := 0; i < len(characters); i++ {
		reg[characters[i]]++
	}
	for i := 0; i < len(document); i++ {
		if _, ok := reg[document[i]]; !ok {
			return false
		}
		reg[document[i]]--
		if reg[document[i]] == 0 {
			delete(reg, document[i])
		}
	}
	return true
}

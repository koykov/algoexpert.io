package first_non_repeating_character

func FirstNonRepeatingCharacter(str string) int {
	reg := make(map[byte]int, len(str))
	for i := 0; i < len(str); i++ {
		reg[str[i]]++
	}
	for i := 0; i < len(str); i++ {
		if reg[str[i]] == 1 {
			return i
		}
	}
	return -1
}

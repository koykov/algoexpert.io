package validate_subsequence

func IsValidSubsequence(array []int, sequence []int) bool {
	la, ls := len(array), len(sequence)
	if la == 0 || ls == 0 || la < ls {
		return false
	}

	for i := 0; i < la; i++ {
		if array[i] == sequence[0] {
			if sequence = sequence[1:]; len(sequence) == 0 {
				return true
			}
		}
	}
	return len(sequence) == 0
}

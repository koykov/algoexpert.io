package validate_subsequence

func IsValidSubsequenceOptimized(array []int, sequence []int) bool {
	la, ls := len(array), len(sequence)
	if la == 0 || ls == 0 || la < ls {
		return false
	}

	// BCE optimization
	_ = sequence[ls-1]
	for i := la - 1; i >= 0; i-- {
		if array[i] == sequence[ls-1] {
			if sequence = sequence[:ls-1]; len(sequence) == 0 {
				return true
			}
			ls--
		}
	}
	return len(sequence) == 0
}

package validate_subsequence

func IsValidSubsequenceOptimized(array []int, sequence []int) bool {
	la, ls := len(array), len(sequence)
	if la == 0 || ls == 0 || la < ls {
		return false
	}

	_ = sequence[ls-1]             // first BCE optimization
	for i := la - 1; i >= 0; i-- { // second BCE optimization (loop from end)
		if array[i] == sequence[ls-1] {
			if sequence = sequence[:ls-1]; len(sequence) == 0 {
				return true
			}
			ls--
		}
	}
	return len(sequence) == 0
}

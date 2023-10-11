package first_duplicate_value

func FirstDuplicateValue(a []int) int {
	if len(a) < 2 {
		return -1
	}
	for i := 1; i <= len(a); i++ {
		n := abs(a[i-1])
		if a[n-1] *= -1; a[n-1] > 0 {
			return n
		}
	}
	return -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

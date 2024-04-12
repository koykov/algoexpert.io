package min_number_of_jumps

func MinNumberOfJumps(a []int) int {
	if len(a) == 1 {
		return 0
	}
	reach, s, j := a[0], a[0], 0
	for i := 1; i < len(a)-1; i++ {
		reach = max(reach, a[i]+i)
		if s = s - 1; s == 0 {
			j++
			s = reach - i
		}
	}
	return j + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

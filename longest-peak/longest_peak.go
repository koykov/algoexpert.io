package longest_peak

func LongestPeak(a []int) int {
	n := len(a)
	if n < 3 {
		return 0
	}

	lp := 0
	mode := true // true - increasing; false - decreasing
	ln := 1
	for i := 1; i < n-1; i++ {
		switch {
		case a[i-1] < a[i] && a[i] > a[i+1]:
			mode = false
			ln += 2
			i += 2
			if i == n-1 {
				ln++
			}
			if i == n-2 && a[i-1] > a[i] {
				ln++
			}
		case a[i-1] < a[i] && mode:
			ln++
		case a[i-1] >= a[i] && mode:
			ln = 1
		case a[i] > a[i+1] && !mode:
			ln++
		case a[i] <= a[i+1] && !mode:
			lp = max(lp, ln)
			ln = 1
			mode = true
		}
	}
	if mode {
		return 0
	}
	return max(lp, ln)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

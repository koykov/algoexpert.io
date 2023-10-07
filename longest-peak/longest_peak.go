package longest_peak

func LongestPeak(a []int) int {
	n := len(a)
	if n < 3 {
		return 0
	}

	l, lm, r, rm := 1, 0, 1, 0
	_ = r
	for i := 0; i < n-1; i++ {
		if i < n-1 {
			switch {
			case a[i+1]-a[i] >= 0:
				l++
			case a[i+1]-a[i] < 0:
				if l >= 2 && (i > 0 && a[i] != a[i-1]) {
					lm = max(lm, l+1)
				}
				l = 1
				// case a[i+1]-a[i] == 0:
				// 	l = 1
			}
		}
		if i < n-1 {
			r1, r2 := a[n-i-1], a[n-i-2]
			switch {
			case r2-r1 >= 0:
				r++
			case r2-r1 < 0:
				if r >= 3 {
					rm = max(rm, r+1)
				}
				r = 1
				// case r2-r1 == 0:
				// 	r = 1
			}
		}
	}
	return max(lm, rm)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

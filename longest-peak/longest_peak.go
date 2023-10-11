package longest_peak

func LongestPeak(a []int) int {
	n := len(a)
	if n < 3 {
		return 0
	}

	var lp, ln int
	for i := 1; i < n-1; i++ {
		if a[i-1] < a[i] && a[i] > a[i+1] {
			ln = 1
			for j := i; j > 0; j-- {
				if a[j] <= a[j-1] {
					break
				}
				ln++
			}
			for j := i; j < n-1; j++ {
				if a[j] <= a[j+1] {
					break
				}
				ln++
			}
			lp = max(lp, ln)
		}
	}
	return lp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package shifted_binary_search

func ShiftedBinarySearch(a []int, t int) int {
	if a[0] == t {
		return 0
	}
	n := len(a)
	lo, hi, m := 0, n-1, 0
	for lo <= hi && lo < n && hi > 0 {
		m = lo + (hi-lo)/2
		if a[m] == t {
			return m
		}
		lsort := a[m] > a[lo]
		if a[m] > t {
			if lsort && a[lo] > t {
				lo = m + 1
			} else {
				hi = m - 1
			}
		} else {
			if !lsort && a[hi] < t {
				hi = m - 1
			} else {
				lo = m + 1
			}
		}
	}
	return -1
}

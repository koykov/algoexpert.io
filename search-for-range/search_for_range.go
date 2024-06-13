package search_for_range

func SearchForRange(a []int, t int) []int {
	r := [2]int{-1, -1}

	lo, hi := 0, len(a)-1
	for lo < hi {
		m := (lo + hi) / 2
		if a[m] >= t {
			hi = m
		} else {
			lo = m + 1
		}
	}

	if a[lo] == t {
		r[0] = lo
		lo, hi = 0, len(a)-1
		for lo < hi {
			m := (lo + hi + 1) / 2
			if a[m] <= t {
				lo = m
			} else {
				hi = m - 1
			}
		}
		r[1] = lo
	}
	return r[:]
}

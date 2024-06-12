package index_equals_value

func IndexEqualsValue(a []int) (r int) {
	r = -1
	n := len(a)
	lo, hi := 0, n-1
	for lo <= hi {
		m := (lo + hi) / 2
		switch {
		case a[m] == m:
			r = m
			hi = m - 1
		case m > a[m]:
			lo = m + 1
		default:
			hi = m - 1
		}
	}
	return
}

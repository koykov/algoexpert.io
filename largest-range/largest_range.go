package largest_range

import "math"

func LargestRange(a []int) []int {
	n := len(a)
	if n == 1 {
		return []int{a[0], a[0]}
	}
	m := make(map[int]bool, n)
	mn, mx := math.MaxInt, math.MinInt
	for i := 0; i < n; i++ {
		m[a[i]] = false
		mn = min(mn, a[i])
		mx = max(mx, a[i])
	}
	var l, r int
	for i := 0; i < n; i++ {
		if m[a[i]] {
			continue
		}

		var rc, lc int
		for lc = a[i] - 1; lc >= mn; {
			if _, ok := m[lc]; !ok {
				break
			}
			m[lc] = true
			lc--
		}
		lc++
		for rc = a[i]; rc <= mx; {
			if _, ok := m[rc]; !ok {
				break
			}
			m[rc] = true
			rc++
		}
		rc--
		if rc-lc > r-l {
			l, r = lc, rc
		}
	}
	return []int{l, r}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package subarray_sort

func SubarraySort(a []int) []int {
	res := [2]int{-1, -1}
	n := len(a)
	l, r := 0, n-1
	for ; l < n-1 && a[l] <= a[l+1]; l++ {
	}
	for ; r > 0 && a[r] >= a[r-1] && r > l; r-- {
	}
	if l >= r {
		return res[:]
	}

	li, ri := l, r
	mn, mx := a[l], a[l]
	for l <= r {
		mn = min(mn, a[l])
		mx = max(mx, a[l])
		l++
	}
	for ; li >= 0 && a[li] > mn; li-- {
	}
	for ; ri < n && a[ri] < mx; ri++ {
	}
	res[0], res[1] = li+1, ri-1
	return res[:]
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

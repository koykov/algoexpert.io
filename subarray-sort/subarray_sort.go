package subarray_sort

import "math"

func SubarraySort(a []int) []int {
	res := [2]int{-1, -1}

	n := len(a)
	if n == 2 {
		if a[0] > a[1] {
			res[0], res[1] = 0, 1
		}
		return res[:]
	}

	l, r := -1, -1
	for i := 1; i < n; i++ {
		if a[i] < a[i-1] {
			l = i
			break
		}
	}
	if l < 0 {
		return res[:]
	}

	rm := a[l]
	for i := l - 1; i < n; i++ {
		if a[i] > rm {
			rm = a[i]
			continue
		}
		r = i
	}
	if r == -1 {
		return res[:]
	}

	mn, mx := math.MaxInt, -math.MaxInt
	for i := l; i <= r; i++ {
		if mn > a[i] {
			mn = a[i]
		}
		if mx < a[i] {
			mx = a[i]
		}
	}

	for i := 0; i < n; i++ {
		if a[i] > mn {
			l = i
			break
		}
	}
	for i := n - 1; i >= 0; i-- {
		if a[i] < mx {
			r = i
			break
		}
	}
	res[0], res[1] = l, r
	return res[:]
}

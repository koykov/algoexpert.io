package maximize_expression

const minInt = -1e9

func MaximizeExpression(arr []int) int {
	r := maxp(arr, 0, 2, 0)
	if r <= minInt {
		return 0
	}
	return r
}

func maxp(a []int, i, k int, state int) (r int) {
	if i >= len(a) {
		return minInt
	}
	if k == 0 {
		return 0
	}
	r = minInt
	if state == 0 {
		t := maxp(a, i+1, k, 1)
		if t != minInt {
			t = a[i] + maxp(a, i+1, k, 1)
		}
		r = max(r, max(maxp(a, i+1, k, 0), t))
	} else {
		t := maxp(a, i+1, k-1, 0)
		if t != minInt {
			t = -a[i] + maxp(a, i+1, k-1, 0)
		}
		r = max(r, max(maxp(a, i+1, k, 1), t))
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

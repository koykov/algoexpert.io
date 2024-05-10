package maximize_expression

import "math"

func MaximizeExpression(arr []int) int {
	n := len(arr)
	if n < 4 {
		return 0
	}
	var mx, mn, a, b, c, d int
	buf := make([]int, n)

	// maximize a
	mx = -math.MaxInt
	for i := 0; i < n; i++ {
		mx = max(arr[i], mx)
		buf[i] = mx
	}
	mx = -math.MaxInt
	for i := 0; i < n-3; i++ {
		if buf[i] > mx {
			mx = buf[i]
			a = i
		}
	}
	clear(buf)

	// minimize b
	mn = math.MaxInt
	for i := a + 1; i < n; i++ {
		mn = min(mn, arr[i])
		buf[i] = mn
	}
	mn = math.MaxInt
	for i := a + 1; i < n-2; i++ {
		if buf[i] < mn {
			mn = buf[i]
			b = i
		}
	}
	clear(buf)

	// maximize c
	mx = -math.MaxInt
	for i := c + 1; i < n; i++ {
		mx = max(arr[i], mx)
		buf[i] = mx
	}
	mx = -math.MaxInt
	for i := b + 1; i < n-1; i++ {
		if buf[i] > mx {
			mx = buf[i]
			c = i
		}
	}
	clear(buf)

	// minimize d
	mn = math.MaxInt
	for i := a + 1; i < n; i++ {
		mn = min(mn, arr[i])
		buf[i] = mn
	}
	mn = math.MaxInt
	for i := c + 1; i < n; i++ {
		if buf[i] < mn {
			mn = buf[i]
			d = i
		}
	}

	return arr[a] - arr[b] + arr[c] - arr[d]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func clear(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
	}
}

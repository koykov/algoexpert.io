package kadane_s_algorithm

import "math"

func KadanesAlgorithm(a []int) int {
	c, n := -math.MaxInt, 0
	for i := 0; i < len(a); i++ {
		n += a[i]
		n = max(n, a[i])
		c = max(c, n)
	}
	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

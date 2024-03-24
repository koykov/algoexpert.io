package max_sum_increasing_subsequence

import "math"

func MaxSumIncreasingSubsequence(a []int) (int, []int) {
	n := len(a)
	buf := make([]int, 0, n*2)
	var s, ms, off int
	ms = -math.MaxInt
	// backward check
	for i := n - 1; i >= 0; i-- {
		c := math.MaxInt
		s = 0
		for j := i; j >= 0; j-- {
			if a[j] < c {
				buf = append(buf, a[j])
				c = a[j]
				s += c
			}
		}
		if s > ms {
			ms = s
			copy(buf[0:], buf[off:])
			off = len(buf) - off
		}
		buf = buf[:off]
	}
	for i := 0; i < len(buf)/2; i++ {
		buf[i], buf[len(buf)-i-1] = buf[len(buf)-i-1], buf[i]
	}

	// forward check
	for i := 0; i < n; i++ {
		c := -math.MaxInt
		s = 0
		for j := i; j < n; j++ {
			if a[j] > c {
				buf = append(buf, a[j])
				c = a[j]
				s += c
			}
		}
		if s > ms {
			ms = s
			copy(buf[0:], buf[off:])
			off = len(buf) - off
		}
		buf = buf[:off]
	}
	return ms, buf
}

package max_sum_increasing_subsequence

import "math"

func MaxSumIncreasingSubsequence(a []int) (int, []int) {
	n := len(a)
	buf, sum := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		buf[i] = math.MaxInt
		sum[i] = a[i]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if a[i] > a[j] {
				if x := sum[j] + a[i]; x > sum[i] {
					sum[i] = x
					buf[i] = j
				}
			}
		}
	}

	x, p := math.MinInt, 0
	for i := 0; i < n; i++ {
		if sum[i] > x {
			x = sum[i]
		}
		if x == sum[i] {
			p = i
		}
	}

	res := sum[:0]
	for p != math.MaxInt {
		res = append(res, a[p])
		p = buf[p]
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return x, res
}

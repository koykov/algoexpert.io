package min_number_of_coins_for_change

import "math"

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	buf := make([]int, n+1)
	for i := 1; i <= n; i++ {
		buf[i] = math.MaxInt / 2
	}
	for i := 0; i < len(denoms); i++ {
		for j := 1; j <= n; j++ {
			if denoms[i] <= j {
				buf[j] = min(1+buf[j-denoms[i]], buf[j])
			}
		}
	}
	if buf[n] == math.MaxInt/2 {
		return -1
	}
	return buf[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

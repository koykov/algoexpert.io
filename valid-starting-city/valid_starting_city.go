package valid_starting_city

import "math"

func ValidStartingCity(distances []int, fuel []int, mpg int) int {
	n, m := len(distances), 0
	mm, mi := math.MaxInt, 0
	for i := 0; i < n; i++ {
		m += fuel[i] * mpg
		m -= distances[i]
		if m < mm {
			mm, mi = m, i
		}
	}
	if mi+1 == n {
		return 0
	}
	return mi + 1
}

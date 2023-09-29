package selection_sort

import "math"

func SelectionSort(a []int) []int {
	for i := len(a) - 1; i > 0; i-- {
		m, mi := -math.MaxInt, 0
		for j := 0; j <= i; j++ {
			if a[j] > m {
				m, mi = a[j], j
			}
		}
		a[mi], a[i] = a[i], a[mi]
	}
	return a
}

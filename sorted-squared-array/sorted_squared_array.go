package sorted_squared_array

import "sort"

func SortedSquaredArray(array []int) []int {
	n := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		n[i] = array[i] * array[i]
	}
	sort.Ints(n)
	return n
}

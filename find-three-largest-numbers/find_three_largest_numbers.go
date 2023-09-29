package find_three_largest_numbers

import "math"

func FindThreeLargestNumbers(array []int) []int {
	r := [3]int{-math.MaxInt, -math.MaxInt, -math.MaxInt}
	for i := 0; i < len(array); i++ {
		if r[0] < array[i] {
			r[0] = array[i]
			if r[0] > r[1] {
				r[0], r[1] = r[1], r[0]
			}
			if r[1] > r[2] {
				r[1], r[2] = r[2], r[1]
			}
		}
	}
	return r[:]
}

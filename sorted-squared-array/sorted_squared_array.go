package sorted_squared_array

func SortedSquaredArray(array []int) []int {
	n := make([]int, len(array))
	if array[0] < 0 && len(array) > 1 {
		h, t := 0, len(array)-1
		i := t
		for ; i >= 0; i-- {
			hq, tq := array[h]*array[h], array[t]*array[t]
			if hq >= tq {
				n[i] = hq
				h++
			} else {
				n[i] = tq
				t--
			}
		}
	} else {
		for i := 0; i < len(array); i++ {
			n[i] = array[i] * array[i]
		}
	}
	return n
}

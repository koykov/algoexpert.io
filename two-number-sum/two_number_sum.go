package two_number_sum

func TwoNumberSum(array []int, target int) []int {
	m := make(map[int]int, len(array))
	for i := 0; i < len(array); i++ {
		m[array[i]] = i
	}
	for i := 0; i < len(array); i++ {
		d := target - array[i]
		if j, ok := m[d]; ok && i != j {
			return []int{array[i], d}
		}
	}
	return []int{}
}

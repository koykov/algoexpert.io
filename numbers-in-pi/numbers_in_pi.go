package numbers_in_pi

import "math"

func NumbersInPi(pi string, numbers []string) int {
	nr := make(map[string]struct{}, len(numbers))
	for i := 0; i < len(numbers); i++ {
		nr[numbers[i]] = struct{}{}
	}
	m := make([]int, len(pi)+1)
	for i := len(pi) - 1; i >= 0; i-- {
		var c string
		a := math.MaxInt64
		for j := i; j < len(pi); j++ {
			c = c + string(pi[j])
			if _, ok := nr[c]; ok {
				a = min(a, m[j+1])
			}
		}
		m[i] = a + 1
	}

	if a := m[0]; a != math.MaxInt {
		return a - 1
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

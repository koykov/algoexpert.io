package next_greater_element

import "math"

func NextGreaterElement(a []int) []int {
	n := len(a)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = -1
	}

	stk := make([]int, 0, len(a))
	mx := -math.MaxInt
	for i := 0; i < len(a); i++ {
		if mx < a[i] {
			mx = a[i]
		}
		if len(stk) == 0 {
			stk = append(stk, i)
			continue
		}
		si := stk[len(stk)-1]
		if a[i] > a[si] {
			for j := len(stk) - 1; j >= 0; j-- {
				r[stk[j]] = a[i]
			}
			stk = stk[:0]
			stk = append(stk, i)
			continue
		} else {
			stk = append(stk, i)
		}
	}
	for i := len(stk) - 1; i >= 0; i-- {
		if a[stk[i]] == mx {
			continue
		}
		r[stk[i]] = r[0]
	}
	return r
}

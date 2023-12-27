package next_greater_element

func NextGreaterElement(a []int) []int {
	n := len(a)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = -1
	}
	stk := make([]int, 0, len(a))
	for i := n - 1; i >= 0; i-- {
		stk = append(stk, a[i])
	}
	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && stk[len(stk)-1] <= a[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			stk = append(stk, a[i])
			continue
		}
		r[i] = stk[len(stk)-1]
		stk = append(stk, a[i])
	}
	return r
}

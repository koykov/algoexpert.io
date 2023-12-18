package sort_stack

func SortStack(stack []int) []int {
	if len(stack) == 0 {
		return stack
	}
	for i := 0; i < len(stack); i++ {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		stack = sort1(stack, n)
	}
	return stack
}

func sort1(stack []int, n int) []int {
	if len(stack) == 0 {
		stack = append(stack, n)
		return stack
	}
	m := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	if n < m {
		n, m = m, n
	}
	stack = sort1(stack, m)
	stack = append(stack, n)
	return stack
}

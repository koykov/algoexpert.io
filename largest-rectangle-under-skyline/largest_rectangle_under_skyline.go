package largest_rectangle_under_skyline

func LargestRectangleUnderSkyline(b []int) int {
	if len(b) == 0 {
		return 0
	}
	if len(b) == 1 {
		return b[0]
	}
	var mx int
	stk := stack{}
	for i := 0; i <= len(b); i++ {
		var c int
		if i < len(b) {
			c = b[i]
		}

		for !stk.empty() && c < b[stk.top()] {
			p := b[stk.top()]
			stk.pop()
			w := i
			if !stk.empty() {
				w = i - stk.top() - 1
			}
			mx = max(mx, w*p)
		}

		stk.push(i)
	}

	return mx
}

type stack []int

func (s *stack) push(v int) {
	*s = append(*s, v)
}

func (s *stack) pop() {
	if len(*s) == 0 {
		return
	}
	*s = (*s)[:len(*s)-1]
}

func (s *stack) top() int {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

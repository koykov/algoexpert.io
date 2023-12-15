package min_max_stack_construction

type entry struct {
	val, min, max int
}

type MinMaxStack struct {
	stack []entry
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{}
}

func (s *MinMaxStack) Peek() int {
	if len(s.stack) == 0 {
		return -1
	}
	return s.stack[len(s.stack)-1].val
}

func (s *MinMaxStack) Pop() int {
	e := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return e.val
}

func (s *MinMaxStack) Push(n int) {
	if len(s.stack) == 0 {
		s.stack = append(s.stack, entry{n, n, n})
	} else {
		min, max := n, n
		top := s.stack[len(s.stack)-1]
		if min > top.min {
			min = top.min
		}
		if max < top.max {
			max = top.max
		}
		s.stack = append(s.stack, entry{n, min, max})
	}
}

func (s *MinMaxStack) GetMin() int {
	return s.stack[len(s.stack)-1].min
}

func (s *MinMaxStack) GetMax() int {
	return s.stack[len(s.stack)-1].max
}

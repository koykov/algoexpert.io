package min_heap_construction

type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h *MinHeap) BuildHeap(a []int) {
	for i := len(a) - 1; i >= 0; i-- {
		h.siftDown(i, len(a)-1)
	}
}

func (h *MinHeap) siftDown(c, e int) {
	l := c*2 + 1
	if l > e {
		return
	}
	r := l + 1
	s := r
	if r > e || (*h)[l] < (*h)[r] {
		s = l
	}
	if (*h)[s] < (*h)[c] {
		(*h)[s], (*h)[c] = (*h)[c], (*h)[s]
		c = s
		h.siftDown(c, e)
	}
}

func (h *MinHeap) siftUp() {
	n := len(*h) - 1
	if n <= 0 {
		return
	}
	p := (n - 1) / 2
	for p >= 0 && (*h)[n] < (*h)[p] {
		(*h)[n], (*h)[p] = (*h)[p], (*h)[n]
		n = p
		p = (n - 1) / 2
	}
}

func (h MinHeap) Peek() int {
	if len(h) > 0 {
		return h[0]
	}
	return -1
}

func (h *MinHeap) Remove() int {
	n := len(*h)
	if n == 0 {
		return -1
	}
	v, idx := (*h)[0], 0
	(*h)[idx] = (*h)[n-1]
	*h = (*h)[:n-1]
	h.siftDown(0, n-1)
	return v
}

func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	h.siftUp()
	return
}

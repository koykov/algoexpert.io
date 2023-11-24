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
		h.siftDown()
	}
}

func (h *MinHeap) siftDown(currentIndex, endIndex int) {
}

func (h *MinHeap) siftUp() {
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
	for {
		l, r, i := 2*idx+1, 2*idx+2, idx
		switch {
		case l < n && (*h)[l] < (*h)[i]:
			i = l
		case r < n && (*h)[r] < (*h)[i]:
			i = r
		case i != idx:
			(*h)[idx], (*h)[i] = (*h)[i], (*h)[idx]
			idx = i
		default:
			goto exit
		}
	}
exit:
	return v
}

func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	n := len(*h) - 1
	for n > 0 && (*h)[(n-1)/2] > (*h)[n] {
		(*h)[n], (*h)[(n-1)/2] = (*h)[(n-1)/2], (*h)[n]
		n = (n - 1) / 2
	}
}

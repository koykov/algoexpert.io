package min_heap_construction

type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h *MinHeap) BuildHeap(array []int) {
}

func (h *MinHeap) siftDown(currentIndex, endIndex int) {
}

func (h *MinHeap) siftUp() {
}

func (h MinHeap) Peek() int {
	return -1
}

func (h *MinHeap) Remove() int {
	return -1
}

func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	n := len(*h) - 1
	for n > 0 && (*h)[(n-1)/2] > (*h)[n] {
		(*h)[n], (*h)[(n-1)/2] = (*h)[(n-1)/2], (*h)[n]
		n = (n - 1) / 2
	}
}

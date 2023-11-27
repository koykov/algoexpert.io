package min_heap_construction

import "testing"

func TestMinHeap(t *testing.T) {
	h := NewMinHeap([]int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41})
	t.Log(h)
	h.Insert(76)
	t.Log(h)
	t.Log(h.Peek())
	h.Remove()
	t.Log(h)
	t.Log(h.Peek())
	h.Remove()
	t.Log(h.Peek())
	h.Insert(87)
	t.Log(h)
}

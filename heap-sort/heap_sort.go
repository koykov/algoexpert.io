package heap_sort

func HeapSort(a []int) []int {
	heapSort(a, len(a))
	return a
}

func heapSort(a []int, ln int) {
	for i := ln/2 - 1; i >= 0; i-- {
		heapify(a, ln, i)
	}
	for i := ln - 1; i > 0; i-- {
		val := a[0]
		a[0] = a[i]
		a[i] = val
		heapify(a, i, 0)
	}
}

func heapify(a []int, ln, idx int) {
	big, left, right := idx, 2*idx+1, 2*idx+2
	if left < ln && a[left] > a[big] {
		big = left
	}
	if right < ln && a[right] > a[big] {
		big = right
	}
	if big != idx {
		val := a[idx]
		a[idx] = a[big]
		a[big] = val
		heapify(a, ln, big)
	}
}

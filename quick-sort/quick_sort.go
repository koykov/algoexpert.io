package quick_sort

func QuickSort(a []int) []int {
	quickSort(a, 0, len(a)-1)
	return a
}

func quickSort(a []int, start, pivot int) {
	var m int
	if start < pivot {
		m = part(a, start, pivot)
		quickSort(a, start, m-1)
		quickSort(a, m+1, pivot)
	}
}

func part(a []int, start, pivot int) int {
	var pv, i int
	pv = a[pivot]
	i = start - 1
	for j := start; j <= pivot-1; j++ {
		if a[j] < pv {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[pivot] = a[pivot], a[i+1]
	return i + 1
}

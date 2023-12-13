package three_number_sort

func ThreeNumberSort(a []int, o []int) []int {
	var off int
	off = walk(a, o[0], off)
	off = walk(a, o[1], off)
	return a
}

func walk(a []int, t, off int) int {
	for i := off; i < len(a); i++ {
		if a[i] == t {
			a[off], a[i] = a[i], a[off]
			off++
		}
	}
	return off
}

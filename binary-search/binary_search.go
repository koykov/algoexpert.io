package binary_search

func BinarySearch(array []int, target int) int {
	return bs(array, target, 0)
}

func bs(a []int, t, off int) (r int) {
	if len(a) == 0 {
		return -1
	}
	if len(a) == 1 {
		if a[0] == t {
			return off
		}
		return -1
	}
	l2 := len(a) / 2
	al, ar := a[:l2], a[l2:]
	if r = bs(al, t, off); r != -1 {
		return r
	}
	if r = bs(ar, t, off+l2); r != -1 {
		return r
	}
	return -1
}

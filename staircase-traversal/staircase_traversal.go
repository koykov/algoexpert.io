package staircase_traversal

func StaircaseTraversal(h int, s int) (r int) {
	if h == 0 {
		r = 1
		return
	}
	for i := 1; i <= s && i <= h; i++ {
		r += StaircaseTraversal(h-i, s)
	}
	return
}

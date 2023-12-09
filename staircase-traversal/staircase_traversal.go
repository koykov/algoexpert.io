package staircase_traversal

func StaircaseTraversal(h int, s int) (r int) {
	if h == 0 || h == 1 {
		return 1
	}
	for i := 1; i <= s; i++ {
		r += StaircaseTraversal(h-i, s)
	}
	return
}

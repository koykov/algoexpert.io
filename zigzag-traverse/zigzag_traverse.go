package zigzag_traverse

func ZigzagTraverse(a [][]int) []int {
	n1, m := len(a), len(a[0])
	r := make([]int, 0, n1*m)
	n := max(n1, m)
	for i := 0; i < n*2; i++ {
		if i%2 == 0 {
			y := i
			for j := 0; j <= i; j++ {
				x := j
				if x < n1 && y < m {
					r = append(r, a[x][y])
				}
				y--
			}
		} else {
			y := 0
			for j := i; j >= 0; j-- {
				x := j
				if x < n1 && y < m {
					r = append(r, a[x][y])
				}
				y++
			}
		}
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

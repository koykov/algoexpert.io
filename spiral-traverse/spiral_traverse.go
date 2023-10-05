package spiral_traverse

func SpiralTraverse(a [][]int) []int {
	x0, y0, x1, y1 := 0, 0, len(a), len(a[0])
	buf := make([]int, 0, x1*y1)
	for {
		if x0 < x1 {
			for i := y0; i < y1; i++ {
				buf = append(buf, a[x0][i])
			}
			x0++
		}
		if y0 < y1 {
			for i := x0; i < x1; i++ {
				buf = append(buf, a[i][y1-1])
			}
			y1--
		}
		if x0 < x1 {
			for i := y1 - 1; i >= y0; i-- {
				buf = append(buf, a[x1-1][i])
			}
			x1--
		}
		if y0 < y1 {
			for i := x1 - 1; i >= x0; i-- {
				buf = append(buf, a[i][x0-1])
			}
			y0++
		}
		if x0 == x1 && y0 == y1 {
			break
		}
	}
	return buf
}

package count_squares

func CountSquares(p [][]int) (c int) {
	reg := make(map[int]bool, len(p))
	var x1, x2, y1, y2, x1_, x2_, y1_, y2_, xm, ym, dx, dy int
	for i := 0; i < len(p); i++ {
		x1, y1 = (p[i][0]+100)*2, (p[i][1]+100)*2
		for j := 0; j < len(p); j++ {
			if i == j {
				continue
			}
			x2, y2 = (p[j][0]+100)*2, (p[j][1]+100)+2
			xm, ym = (x1+x2)/2, (y1+y2)/2
			dx, dy = x1-xm, y1-ym
			x1_, y1_ = xm-dy, ym+dx
			x2_, y2_ = xm+dy, ym-dx

			k1, k2 := x1_<<32+y1_, x2_<<32+y2_
			if reg[k1] && reg[k2] {
				c++
			}
		}
		reg[x1<<32+y1] = true
	}
	return
}

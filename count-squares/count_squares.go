package count_squares

func CountSquares(p [][]int) (c int) {
	reg := make(map[float64]bool, len(p))
	for i := 0; i < len(p); i++ {
		reg[comb(float64(p[i][0]), float64(p[i][1]))] = true
	}
	var x1, x2, y1, y2, x1_, x2_, y1_, y2_, mx, my, dx, dy float64
	for i := 0; i < len(p); i++ {
		x1, y1 = float64(p[i][0]), float64(p[i][1])
		for j := i + 1; j < len(p); j++ {
			x2, y2 = float64(p[j][0]), float64(p[j][1])
			mx, my = (x1+x2)/2, (y1+y2)/2
			dx, dy = x2-mx, y2-my
			if dx == 0 && dy == 0 {
				continue
			}
			x1_, y1_ = mx+dy, my-dx
			x2_, y2_ = mx-dy, my+dx

			k1, k2 := comb(x1_, y1_), comb(x2_, y2_)
			if reg[k1] || reg[k2] {
				print(" ")
			}
			if reg[k1] && reg[k2] {
				c++
			}
		}
	}
	c /= 2
	return
}

func comb(a, b float64) float64 { return a*1e6 + b }

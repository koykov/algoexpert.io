package minimum_passes_of_matrix

import "strconv"

func MinimumPassesOfMatrix(mx [][]int) (c int) {
	n, m := len(mx), len(mx[0])
	buf := make([]int, n*m)
	mx1 := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		mx1 = append(mx1, buf[i*m:i*m+m])
	}
	mirror(mx, mx1)

	for {
		var neg, pass bool
		for i := 0; i < len(mx); i++ {
			for j := 0; j < len(mx[i]); j++ {
				if mx[i][j] < 0 {
					pass1 := (j > 0 && mx[i][j-1] > 0) ||
						(i+1 < len(mx) && mx[i+1][j] > 0) ||
						(j+1 < len(mx[i]) && mx[i][j+1] > 0) ||
						(i > 0 && mx[i-1][j] > 0)
					if pass1 {
						pass = true
						mx1[i][j] *= -1
					}
				}
				neg = neg || mx[i][j] < 0
			}
		}
		mirror(mx1, mx)
		c++
		if !neg {
			break
		}
		if !pass && neg {
			c = 0
			break
		}
	}
	return c - 1
}

func mirror(a, b [][]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			b[i][j] = a[i][j]
		}
	}
}

func mprint(m [][]int) (r string) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			r += strconv.Itoa(m[i][j]) + " "
		}
		r += "\n"
	}
	return
}

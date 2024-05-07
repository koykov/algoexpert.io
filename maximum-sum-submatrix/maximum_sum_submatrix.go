package maximum_sum_submatrix

import "math"

func MaximumSumSubmatrix(m [][]int, sz int) int {
	s := make([][]int, len(m))
	var p int
	for i := 0; i < len(m); i++ {
		s[i] = make([]int, len(m[i]))
		s[i][0] = p + m[i][0]
		p = s[i][0]
	}
	p = 0
	for i := 0; i < len(m[0]); i++ {
		s[0][i] = p + m[0][i]
		p = s[0][i]
	}
	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[i]); j++ {
			s[i][j] = m[i][j] + s[i-1][j] + s[i][j-1] - s[i-1][j-1]
		}
	}
	mx := -math.MaxInt
	for i := sz - 1; i < len(m); i++ {
		for j := sz - 1; j < len(m[i]); j++ {
			var l, r, lr int
			if i >= sz {
				l = s[i-sz][j]
			}
			if j >= sz {
				r = s[i][j-sz]
			}
			if i >= sz && j >= sz {
				lr = s[i-sz][j-sz]
			}
			if x := s[i][j] - l - r + lr; x > mx {
				mx = x
			}
		}
	}
	return mx
}

func prnt(m [][]int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			print(m[i][j], " ")
		}
		println()
	}
	println()
}

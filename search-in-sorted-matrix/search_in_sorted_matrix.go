package search_in_sorted_matrix

func SearchInSortedMatrix(m [][]int, t int) []int {
	r := [2]int{-1, -1}

	c := -1
	for i := 1; i < len(m); i++ {
		switch {
		case m[i][0] == t:
			r[0], r[1] = i, 0
			return r[:]
		case m[i][0] > t && m[i-1][0] < t:
			c = i - 1
			break
		}
	}
	if c >= 0 {
		for i := 0; i < len(m[c]); i++ {
			if m[c][i] == t {
				r[0], r[1] = c, i
				return r[:]
			}
		}
	}

	for i := 1; i < len(m[0]); i++ {
		switch {
		case m[0][i] == t:
			r[0], r[1] = 0, i
			return r[:]
		case m[0][i] > t && m[0][i-1] < t:
			c = i - 1
			break
		}
	}
	if c >= 0 {
		for i := 0; i < len(m); i++ {
			if m[i][c] == t {
				r[0], r[1] = i, c
				return r[:]
			}
		}
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == t {
				r[0], r[1] = i, j
				goto exit
			}
		}
	}
exit:
	return r[:]
}

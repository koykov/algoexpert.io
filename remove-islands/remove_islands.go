package remove_islands

import "strconv"

func RemoveIslands(m [][]int) [][]int {
	redge(m, 1, 2)
	for i := 1; i < len(m)-1; i++ {
		for j := 1; j < len(m[i])-1; j++ {
			if m[i][j] == 1 {
				m[i][j] = 0
			}
		}
	}
	redge(m, 2, 1)
	return m
}

func redge(m [][]int, o, n int) {
	for i := 0; i < len(m); i++ {
		if m[i][0] == o {
			rmark(m, i, 0, o, n)
		}
		if m[i][len(m[i])-1] == o {
			rmark(m, i, len(m[i])-1, o, n)
		}
	}
	for i := 0; i < len(m[0]); i++ {
		if m[0][i] == o {
			rmark(m, 0, i, o, n)
		}
		if m[len(m)-1][i] == o {
			rmark(m, len(m)-1, i, o, n)
		}
	}
}

func rmark(m [][]int, x, y, o, n int) {
	m[x][y] = n
	if y > 0 && m[x][y-1] == o {
		rmark(m, x, y-1, o, n)
	}
	if x+1 < len(m) && m[x+1][y] == o {
		rmark(m, x+1, y, o, n)
	}
	if y+1 < len(m) && m[x][y+1] == o {
		rmark(m, x, y+1, o, n)
	}
	if x > 0 && m[x-1][y] == o {
		rmark(m, x-1, y, o, n)
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

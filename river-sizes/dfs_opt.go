package river_sizes

func RiverSizesOptimized(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	buf := make([]int, 0, 16)
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if matrix[x][y] == 1 {
				c := dfs1(matrix, x, y)
				buf = append(buf, c)
			}
		}
	}
	return buf
}

func dfs1(matrix [][]int, x, y int) int {
	matrix[x][y] = 0
	c := 1
	if x > 0 && matrix[x-1][y] == 1 {
		c += dfs1(matrix, x-1, y)
	}
	if y+1 < len(matrix[x]) && matrix[x][y+1] == 1 {
		c += dfs1(matrix, x, y+1)
	}
	if x+1 < len(matrix) && matrix[x+1][y] == 1 {
		c += dfs1(matrix, x+1, y)
	}
	if y > 0 && matrix[x][y-1] == 1 {
		c += dfs1(matrix, x, y-1)
	}
	return c
}

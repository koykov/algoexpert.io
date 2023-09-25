package river_sizes

func RiverSizes(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	aux, auxb := make([][]bool, n), make([]bool, n*m)
	for i := 0; i < n; i++ {
		aux[i] = auxb[i*m : i*m+m]
	}
	buf := make([]int, 0, 16)
	for x := 0; x < n; x++ {
		for y := 0; y < len(matrix[x]); y++ {
			if aux[x][y] || matrix[x][y] == 0 {
				aux[x][y] = true
				continue
			}
			c := dfs(matrix, aux, x, y, 0)
			buf = append(buf, c)
		}
	}
	return buf
}

func dfs(matrix [][]int, aux [][]bool, x, y, depth int) int {
	aux[x][y] = true
	c := 1
	if x > 0 && !aux[x-1][y] && matrix[x-1][y] == 1 {
		c += dfs(matrix, aux, x-1, y, depth+1)
	}
	if y+1 < len(matrix[x]) && !aux[x][y+1] && matrix[x][y+1] == 1 {
		c += dfs(matrix, aux, x, y+1, depth+1)
	}
	if x+1 < len(matrix) && !aux[x+1][y] && matrix[x+1][y] == 1 {
		c += dfs(matrix, aux, x+1, y, depth+1)
	}
	if y > 0 && !aux[x][y-1] && matrix[x][y-1] == 1 {
		c += dfs(matrix, aux, x, y-1, depth+1)
	}
	return c
}

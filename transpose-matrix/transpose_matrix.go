package transpose_matrix

func TransposeMatrix(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	buf := make([]int, n*m)
	t := make([][]int, m)
	for i := 0; i < m; i++ {
		t[i] = buf[i*n : i*n+n]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			t[j][i] = matrix[i][j]
		}
	}

	return t
}

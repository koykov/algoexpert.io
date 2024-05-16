package maximize_expression

func MaximizeExpression(arr []int) int {
	n := len(arr)
	if n < 4 {
		return 0
	}
	var m [5][]int
	for i := 0; i < 5; i++ {
		m[i] = make([]int, n+1)
	}

	sign := 1
	for i := 1; i < 5; i++ {
		for j := i; j < n+1; j++ {
			if i == j {
				m[i][j] = m[i-1][j-1] + sign*arr[j-1]
			} else {
				m[i][j] = max(m[i][j-1], m[i-1][j-1]+sign*arr[j-1])
			}
		}
		sign *= -1
	}
	return m[4][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

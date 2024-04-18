package knapsack_problem

func KnapsackProblem(itm [][]int, c int) []interface{} {
	var a [][]int
	for i := 0; i <= len(itm); i++ {
		a = append(a, make([]int, c+1))
	}
	for i := 1; i <= len(itm); i++ {
		for j := 1; j <= c; j++ {
			p, w := itm[i-1][0], itm[i-1][1]
			if j >= w {
				a[i][j] = max(a[i-1][j], a[i-1][j-w]+p)
			} else {
				a[i][j] = a[i-1][j]
			}
		}
	}
	return []interface{}{
		a[len(a)-1][c],
		walk1(a, itm),
	}
}

func walk1(a, itm [][]int) (r []int) {
	r = []int{}
	i, j := len(a)-1, len(a[0])-1
	for i > 0 {
		if a[i][j] > a[i-1][j] {
			c := i - 1
			w := itm[c][1]
			r = append(r, c)
			j = j - w
		}
		i--
	}
	for i = 0; i < len(r)/2; i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

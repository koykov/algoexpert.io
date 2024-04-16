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
	buf, w := walk(nil, a, itm, len(itm), c)
	return []interface{}{
		w,
		buf,
	}
}

func walk(dst []int, a, itm [][]int, k, s int) ([]int, int) {
	var wo int
	if a[k][s] == 0 {
		return dst, wo
	}
	if a[k-1][s] == a[k][s] {
		dst, wo = walk(dst, a, itm, k-1, s)
	} else {
		var wi int
		w := itm[k-1][1]
		dst, wi = walk(dst, a, itm, k-1, s-w)
		dst = append(dst, k-1)
		wo = w + wi
	}
	return dst, wo
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func print_(a [][]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			print(a[i][j], " ")
		}
		print("\n")
	}
}

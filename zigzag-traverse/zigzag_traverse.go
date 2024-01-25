package zigzag_traverse

import "fmt"

func ZigzagTraverse(a [][]int) []int {
	n, m := len(a), len(a[0])
	r := make([]int, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("[%d, %d] %d\n", i, j, a[i][j])
		}
	}
	return r
}

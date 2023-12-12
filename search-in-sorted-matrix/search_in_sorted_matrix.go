package search_in_sorted_matrix

func SearchInSortedMatrix(m [][]int, t int) []int {
	r := [2]int{-1, -1}

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

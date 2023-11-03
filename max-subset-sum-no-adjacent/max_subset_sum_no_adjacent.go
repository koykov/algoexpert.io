package max_subset_sum_no_adjacent

func MaxSubsetSumNoAdjacent(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}
	p2 := a[0]
	p1 := max(a[1], p2)
	m := p1
	for i := 2; i < len(a); i++ {
		m = max(p1, p2+a[i])
		p1, p2 = m, p1
	}
	return max(p1, p2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

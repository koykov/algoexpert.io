package four_number_sum

func FourNumberSum(a []int, t int) [][]int {
	r := make([][]int, 0)
	uniq := make(map[int][][2]int)
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			d := t - a[i] - a[j]
			if pp, ok := uniq[d]; ok {
				for k := 0; k < len(pp); k++ {
					r = append(r, []int{pp[k][0], pp[k][1], a[i], a[j]})
				}
			}
		}
		for k := 0; k < i; k++ {
			s := a[i] + a[k]
			uniq[s] = append(uniq[s], [2]int{a[i], a[k]})
		}
	}
	return r
}

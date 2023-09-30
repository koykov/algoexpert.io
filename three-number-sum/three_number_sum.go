package three_number_sum

import "sort"

func ThreeNumberSum(a []int, t int) (res [][]int) {
	res = make([][]int, 0, 16)
	sort.Ints(a)
	for i := 0; i < len(a); i++ {
		l, r := i, len(a)-1
		for l < r {
			if i == l || l == r {
				l++
				continue
			}
			s := a[i] + a[l] + a[r]
			switch {
			case s < t:
				l++
			case s > t:
				r--
			case s == t:
				res = append(res, []int{a[i], a[l], a[r]})
				l++
			}
		}
	}
	return
}

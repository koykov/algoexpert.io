package smallest_difference

import "sort"

func SmallestDifference(a1, a2 []int) (r []int) {
	sort.Ints(a1)
	sort.Ints(a2)
	r = []int{a1[0], a2[len(a2)-1]}
	x1, x2 := 0, 0
	abs := func(a, b int) (r int) {
		if r = a - b; r < 0 {
			r = -r
		}
		return
	}
	for {
		if a1[x1]-a2[x2] == 0 {
			r[0], r[1] = a1[x1], a2[x2]
			return
		}
		if abs(a1[x1], a2[x2]) < abs(r[0], r[1]) {
			r[0], r[1] = a1[x1], a2[x2]
		}
		if a1[x1] < a2[x2] {
			x1++
		} else {
			x2++
		}
		if x1 == len(a1) || x2 == len(a2) {
			return
		}
	}
}

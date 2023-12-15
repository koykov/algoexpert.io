package sunset_views

import "sort"

func SunsetViews(b []int, d string) (r []int) {
	r = []int{}
	l := -1
	switch d {
	case "EAST":
		for i := len(b) - 1; i >= 0; i-- {
			if b[i] > l {
				l = b[i]
				r = append(r, i)
			}
		}
	case "WEST":
		for i := 0; i < len(b); i++ {
			if b[i] > l {
				l = b[i]
				r = append(r, i)
			}
		}
	}
	sort.Ints(r)
	return
}

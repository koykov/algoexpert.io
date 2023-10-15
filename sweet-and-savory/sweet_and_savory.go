package sweet_and_savory

import (
	"math"
	"sort"
)

func SweetAndSavory(d []int, t int) []int {
	sort.Ints(d)
	e := -1
	for i := 0; i < len(d); i++ {
		if d[i] > 0 {
			e = i - 1
			break
		}
	}
	if e < 0 {
		return []int{0, 0}
	}

	sw, sv, pdiff := 0, 0, -math.MaxInt
	for l := e; l >= 0; l-- {
		for r := e + 1; r < len(d); r++ {
			diff := d[l] + d[r]
			if diff <= t {
				if diff == t {
					return []int{d[l], d[r]}
				}
				if diff > pdiff {
					sw, sv = d[l], d[r]
					pdiff = diff
				}
			} else {
				break
			}
		}
	}
	return []int{sw, sv}
}

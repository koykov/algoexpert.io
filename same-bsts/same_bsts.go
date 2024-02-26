package same_bsts

import "sort"

func SameBsts(a1, a2 []int) bool {
	if len(a1) != len(a2) || a1[0] != a2[0] {
		return false
	}
	sort.Ints(a1)
	sort.Ints(a2)
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
		if i > 0 && a1[i] == a1[i-1] {
			return false
		}
	}
	return true
}

package same_bsts

func SameBsts(a1, a2 []int) bool {
	if len(a1) == 0 && len(a2) == 0 {
		return true
	}
	if len(a1) != len(a2) || a1[0] != a2[0] {
		return false
	}
	buf1, buf2 := make([]int, 0, len(a1)), make([]int, 0, len(a2))
	for i := 1; i < len(a1); i++ {
		if a1[i] < a1[0] {
			buf1 = append(buf1, a1[i])
		}
	}
	for i := 1; i < len(a2); i++ {
		if a2[i] < a2[0] {
			buf2 = append(buf2, a2[i])
		}
	}
	eq1 := SameBsts(buf1, buf2)
	buf1, buf2 = buf1[:0], buf2[:0]
	for i := 1; i < len(a1); i++ {
		if a1[i] >= a1[0] {
			buf1 = append(buf1, a1[i])
		}
	}
	for i := 1; i < len(a2); i++ {
		if a2[i] >= a2[0] {
			buf2 = append(buf2, a2[i])
		}
	}
	eq2 := SameBsts(buf1, buf2)
	return eq1 && eq2
}

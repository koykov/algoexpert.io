package permutations

func GetPermutations(a []int) [][]int {
	r := [][]int{}
	if len(a) == 0 {
		return r
	}
	return permut(r, nil, a)
}

func permut(dst [][]int, buf, a []int) [][]int {
	if len(a) == 0 {
		dst = append(dst, buf)
		return dst
	}
	for i := 0; i < len(a); i++ {
		b := make([]int, 0, len(a))
		b = append(b, a[:i]...)
		b = append(b, a[i+1:]...)
		p := make([]int, 0, len(a))
		p = append(p, buf...)
		p = append(p, a[i])
		dst = permut(dst, p, b)
	}
	return dst
}

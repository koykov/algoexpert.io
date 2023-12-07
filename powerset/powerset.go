package powerset

func Powerset(a []int) [][]int {
	switch {
	case len(a) == 0:
		return [][]int{{}}
	case len(a) == 1:
		return [][]int{{}, a}
	}
	var r [][]int
	cpy := append([]int(nil), a...)
	cpy = cpy[1:]
	t := Powerset(cpy)
	for i := 0; i < len(t); i++ {
		r = append(r, t[i])
		t1 := append([]int(nil), a[0])
		t1 = append(t1, t[i]...)
		r = append(r, t1)
	}
	return r
}

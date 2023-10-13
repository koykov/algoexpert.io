package majority_element

func MajorityElement(a []int) (r int) {
	r = -1
	var n, c int
	for i := 0; i < len(a); i++ {
		if c == 0 {
			n, c = a[i], 1
			continue
		}
		if a[i] == n {
			c++
		} else {
			c--
		}
	}
	if c > 0 {
		r = n
	}
	return
}

package missingNumbers

func MissingNumbers(a []int) []int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	a = append(a, 1, 1)
	r := make([]int, 0, 2)
	for i := 0; i < len(a); i++ {
		a[abs(a[i])-1] *= -1
	}
	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			r = append(r, i+1)
		}
	}
	return r
}

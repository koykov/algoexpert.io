package single_cycle_check

func HasSingleCycle(a []int) bool {
	n := len(a)
	var c, i int
	for {
		c++
		i = (i + a[i]) % n
		i = (i + n) % n
		if i == 0 {
			return c == n
		}
		if c == n {
			return false
		}
	}
}

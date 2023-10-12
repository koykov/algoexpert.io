package best_seat

func BestSeat(a []int) int {
	var m, l, r, o int
	o = -1
	for i := 1; i < len(a); i++ {
		if a[i] == 0 {
			if a[i-1] == 1 {
				l = i
				r = l
			}
			r++
		}
		if a[i] == 1 && a[i-1] == 0 {
			m1 := r - l
			if m1 > m {
				m = m1
				o = l + m/2
				if m%2 == 0 {
					o--
				}
			}
		}
	}
	return o
}

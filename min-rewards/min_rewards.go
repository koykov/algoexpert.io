package min_rewards

func MinRewards(s []int) int {
	n := len(s)
	if n == 1 {
		return 1
	}
	var r int
	w := make([]int, n)
	for i := 0; i < n; i++ {
		if inflectLow(s, i) {
			w[i] = 1
			r = 2
			if i > 0 && w[i-1] < r {
				for j := i - 1; j >= 0; j-- {
					w[j] = r
					r++
					if inflectHigh(s, j) {
						break
					}
				}
			}

			r = 2
			for j := i + 1; j < n; j++ {
				w[j] = r
				r++
				i++
				if inflectHigh(s, j) {
					break
				}
			}
		}
	}
	r = 0
	for i := 0; i < n; i++ {
		r += w[i]
	}
	return r
}

func inflectLow(s []int, i int) bool {
	switch {
	case i == 0:
		return s[i+1] > s[i]
	case i == len(s)-1:
		return s[i-1] > s[i]
	default:
		return s[i-1] > s[i] && s[i] < s[i+1]
	}
}

func inflectHigh(s []int, i int) bool {
	switch {
	case i == 0:
		return s[i+1] < s[i]
	case i == len(s)-1:
		return s[i-1] < s[i]
	default:
		return s[i-1] < s[i] && s[i] > s[i+1]
	}
}

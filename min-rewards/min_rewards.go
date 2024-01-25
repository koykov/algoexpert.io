package min_rewards

func MinRewards(s []int) int {
	n := len(s)
	if n == 1 {
		return 1
	}
	w := make([]int, n)
	off := -1
	for i := 0; i < n; i++ {
		if inflectLow(s, i) {
			r := 2
			w[i] = 1
			for j := i - 1; j > off; j-- {
				w[j] = r
				r++
			}
			off = i

			r = 2
			for j := i + 1; j < n; j++ {
				w[j] = r
				r++
				off++
				i++
				if inflectHigh(s, j) {
					break
				}
			}
		}
	}
	var r int
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

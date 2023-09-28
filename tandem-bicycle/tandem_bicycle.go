package tandem_bicycle

import "sort"

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) (r int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sort.Ints(redShirtSpeeds)
	sort.Ints(blueShirtSpeeds)
	for i := 0; i < len(redShirtSpeeds); i++ {
		if fastest {
			r += max(redShirtSpeeds[i], blueShirtSpeeds[len(blueShirtSpeeds)-i-1])
		} else {
			r += max(redShirtSpeeds[i], blueShirtSpeeds[i])
		}
	}
	return
}

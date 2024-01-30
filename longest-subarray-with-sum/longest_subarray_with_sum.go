package longest_subarray_with_sum

func LongestSubarrayWithSum(a []int, t int) []int {
	var lo, hi, l, h, s int
	for i := 0; i < len(a); i++ {
		if s < t {
			hi = i
			s += a[i]
		}
		switch {
		case s < t:
			continue
		case s == t:
			if hi-lo > h-l {
				l, h = lo, hi
			}
			s -= a[lo]
			lo++
		case s > t:
			s -= a[lo]
			lo++
			i--
		}
	}
	r := [2]int{l, h}
	return r[:]
}

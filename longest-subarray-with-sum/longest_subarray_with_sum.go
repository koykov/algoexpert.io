package longest_subarray_with_sum

func LongestSubarrayWithSum(a []int, t int) []int {
	var lo, hi, rlo, rhi, s int
	for i := 0; i < len(a); i++ {
		switch {
		case s < t:
			hi = i
			s += a[i]
		case s == t:
			if hi-lo > rhi-rlo {
				rlo, rhi = lo, hi
			}
		case s > t:
			s -= a[lo]
			lo++
			i--
		}
	}
	r := [2]int{rlo, rhi}
	return r[:]
}

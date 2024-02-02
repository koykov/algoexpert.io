package longest_subarray_with_sum

func LongestSubarrayWithSum(a []int, t int) []int {
	var lo, hi, lo_, hi_, s int
	for hi < len(a) {
		s += a[hi]
		for ; s > t && lo < hi; lo++ {
			s -= a[lo]
		}
		if s == t && hi-lo >= hi_-lo_ {
			lo_, hi_ = lo, hi
		}
		hi++
	}
	if s != t && lo_ == 0 && hi_ == 0 {
		return []int{}
	}
	r := [2]int{lo_, hi_}
	return r[:]
}

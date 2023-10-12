package zero_sum_subarray

func ZeroSumSubarray(a []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			return true
		}
		var s int
		for j := i; j >= 0; j-- {
			if s += a[j]; s == 0 {
				return true
			}
		}
	}
	return false
}

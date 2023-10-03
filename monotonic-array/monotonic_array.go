package monotonic_array

func IsMonotonic(a []int) bool {
	if len(a) < 2 {
		return true
	}
	inc, dec := true, true
	for i := 0; i < len(a)-1; i++ {
		inc = inc && a[i]-a[i+1] <= 0
		dec = dec && a[i]-a[i+1] >= 0
	}
	return inc || dec
}

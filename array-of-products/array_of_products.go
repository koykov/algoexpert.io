package array_of_products

func ArrayOfProducts(a []int) []int {
	buf := make([]int, len(a))
	t := 1
	for i := 0; i < len(a); i++ {
		buf[i] = t
		t *= a[i]
	}
	t = 1
	for i := len(a) - 1; i >= 0; i-- {
		buf[i] *= t
		t = t * a[i]
	}

	return buf
}

// func ArrayOfProducts(a []int) []int {
// 	t := a[0]
// 	for i := 1; i < len(a); i++ {
// 		t *= a[i]
// 	}
// 	sign := 1
// 	if t < 0 {
// 		sign = -1
// 	}
// 	buf := make([]int, len(a))
// 	if t == 0 {
// 		return buf
// 	}
//
// 	for i := 0; i < len(a); i++ {
// 		buf[i] = sign * int(math.Ceil(math.Pow(10, math.Log10(math.Abs(float64(t)))-math.Log10(math.Abs(float64(a[i]))))))
// 	}
//
// 	return buf
// }

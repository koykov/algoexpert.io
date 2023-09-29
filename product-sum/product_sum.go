package product_sum

type SpecialArray []interface{}

func ProductSum(array []interface{}) int {
	return dive(array, 1)
}

func dive(a []interface{}, depth int) (r int) {
	for i := 0; i < len(a); i++ {
		switch x := a[i].(type) {
		case int:
			r += x
		case SpecialArray:
			r1 := dive(x, depth+1)
			r += (depth + 1) * r1
		}
	}
	return
}

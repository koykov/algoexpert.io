package index_equals_value

func IndexEqualsValue(a []int) int {
	n := len(a)
	m := n / 2
	switch {
	case a[m] == m:
		return m
	case a[m] > m:
		return IndexEqualsValue(a[:m])
	case a[m] < m:
		return IndexEqualsValue(a[m:])
	}
	return -1
}

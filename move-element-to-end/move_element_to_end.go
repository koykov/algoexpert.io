package move_element_to_end

func MoveElementToEnd(a []int, x int) []int {
	var p int
	for i := 0; i < len(a); i++ {
		if a[i] != x {
			a[p], a[i] = a[i], a[p]
			p++
		}
	}
	return a
}

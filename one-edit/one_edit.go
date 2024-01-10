package one_edit

func OneEdit(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	var d int
	if len(s1) == len(s2) {
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				d++
			}
		}
	} else {
		m, n := s1, s2
		if len(s2) > len(s1) {
			m, n = s2, s1
		}
		if len(m)-len(n) > 1 {
			return false
		}
		var off int
		for i := 0; i < len(m); i++ {
			if off == len(n) || m[i] != n[off] {
				d++
				continue
			}
			off++
		}
	}
	return d == 1
}

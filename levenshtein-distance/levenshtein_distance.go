package levenshtein_distance

func LevenshteinDistance(a, b string) int {
	m := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		m[i] = make([]int, len(b)+1)
	}
	for i := 0; i <= len(a); i++ {
		m[i][0] = i
	}
	for i := 0; i <= len(b); i++ {
		m[0][i] = i
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] != b[j-1] {
				m[i][j] = min3(m[i-1][j], m[i][j-1], m[i-1][j-1]) + 1
			} else {
				m[i][j] = m[i-1][j-1]
			}
		}
	}
	return m[len(a)][len(b)]
}

func min3(a, b, c int) (x int) {
	if a < b {
		x = a
	} else {
		x = b
	}
	if c < x {
		x = c
	}
	return
}

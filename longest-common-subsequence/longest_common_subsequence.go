package longest_common_subsequence

func LongestCommonSubsequence(s1 string, s2 string) []byte {
	m, n := len(s1), len(s2)
	var x [][][]byte
	for i := 0; i < n+1; i++ {
		x = append(x, make([][]byte, m+1))
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[j-1] == s2[i-1] {
				x[i][j] = append(x[i][j], x[i-1][j-1]...)
				x[i][j] = append(x[i][j], s2[i-1])
			} else {
				if len(x[i][j-1]) >= len(x[i-1][j]) {
					x[i][j] = x[i][j-1]
				} else {
					x[i][j] = x[i-1][j]
				}
			}
		}
	}
	return x[n][m]
}

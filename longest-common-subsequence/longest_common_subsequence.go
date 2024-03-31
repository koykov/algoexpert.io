package longest_common_subsequence

func LongestCommonSubsequence(s1 string, s2 string) []byte {
	n, m := len(s1), len(s2)
	var x [][]int
	for i := 0; i < m; i++ {
		x = append(x, make([]int, m))
		if i > 0 {
			copy(x[i], x[i-1])
		}
		var inc bool
		for j := 0; j < n; j++ {
			if s2[i] == s1[j] {
				inc = true
			}
			if inc {
				x[i][j]++
			}
		}
	}

	var buf []byte
	var p int
	for i := 0; i < n; i++ {
		if x[m-1][i] != p {
			p = x[m-1][i]
			buf = append(buf, s1[i])
		}
	}
	return buf
}

func print_(x [][]int) {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			print(x[i][j], " ")
		}
		print("\n")
	}
}

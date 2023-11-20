package two_colorable

func TwoColorable(g [][]int) (ok bool) {
	a := make(map[int]int, 100)
	for i := 0; i < len(g); i++ {
		if _, ok_ := a[i]; !ok_ {
			ok = dfs(i, 1, g, a)
		}
		if !ok {
			return
		}
	}
	return true
}

func dfs(v, c int, g [][]int, a map[int]int) bool {
	if c1, ok := a[v]; ok {
		return c1 == c
	}
	a[v] = c
	for i := 0; i < len(g[v]); i++ {
		u := g[v][i]
		if !dfs(u, -c, g, a) {
			return false
		}
	}
	return true
}

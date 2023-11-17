package two_colorable

func TwoColorable(g [][]int) (ok bool) {
	a := make(map[int]int, 100)
	for i := 0; i < len(g); i++ {
		if _, ok_ := a[i]; !ok_ {
			dfs(i, 1, g, a, &ok)
		}
	}
	return
}

func dfs(v, c int, g [][]int, a map[int]int, r *bool) {
	a[v] = c
	for i := 0; i < len(g[v]); i++ {
		u := g[v][i]
		if _, ok := a[u]; !ok {
			dfs(u, -c, g, a, r)
		} else if a[u] != -c {
			*r = true
		}
	}
}

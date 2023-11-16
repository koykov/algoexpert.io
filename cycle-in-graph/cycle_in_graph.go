package cycle_in_graph

func CycleInGraph(g [][]int) (ok bool) {
	reg := make(map[int]byte, 1000)
	for i := 0; i < len(g); i++ {
		if _, ok_ := reg[i]; !ok_ {
			dfs(i, g, reg, &ok)
		}
		if ok {
			break
		}
	}
	return
}

func dfs(v int, g [][]int, reg map[int]byte, ok *bool) {
	reg[v] = 1
	for i := 0; i < len(g[v]); i++ {
		x := g[v][i]
		if reg[v] == 0 {
			dfs(x, g, reg, ok)
		} else if reg[x] == 1 {
			*ok = true
		}
	}
}

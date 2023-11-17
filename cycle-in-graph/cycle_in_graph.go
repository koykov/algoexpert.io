package cycle_in_graph

// CycleInGraph implements topological sort approach.
func CycleInGraph(g [][]int) (ok bool) {
	w := make([]int, 100)
	buf := make([]int, 0, 100)
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			w[g[i][j]]++
		}
	}
	for i := 0; i < len(g); i++ {
		if w[i] == 0 {
			buf = append(buf, i)
		}
	}
	var c int
	for {
		if len(buf) == 0 {
			break
		}
		n := buf[0]
		buf = append(buf[:0], buf[1:]...)
		c++
		for i := 0; i < len(g[n]); i++ {
			n1 := g[n][i]
			w[n1]--
			if w[n1] == 0 {
				buf = append(buf, n1)
			}
		}
	}
	return c != len(g)
}

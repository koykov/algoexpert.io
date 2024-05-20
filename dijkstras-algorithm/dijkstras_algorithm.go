package dijkstras_algorithm

import "math"

func DijkstrasAlgorithm(start int, edges [][][]int) []int {
	buf := make([]int, len(edges))
	v := make([]bool, len(edges))
	for i := 0; i < len(edges); i++ {
		buf[i] = math.MaxInt
	}
	buf[start] = 0
	v[start] = true
	buf = bfs(buf, start, 0, edges, v)
	for i := 0; i < len(buf); i++ {
		if buf[i] == math.MaxInt {
			buf[i] = -1
		}
	}
	return buf
}

func bfs(dst []int, s, val int, e [][][]int, v []bool) []int {
	ce := e[s]
	for i := 0; i < len(ce); i++ {
		c := ce[i]
		dst[c[0]] = min(dst[c[0]], val+c[1])
	}
	for i := 0; i < len(ce); i++ {
		c := ce[i]
		if v[c[0]] {
			continue
		}
		v[c[0]] = true
		dst = bfs(dst, c[0], dst[c[0]], e, v)
	}
	return dst
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

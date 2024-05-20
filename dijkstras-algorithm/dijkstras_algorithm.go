package dijkstras_algorithm

import "math"

func DijkstrasAlgorithm(start int, edges [][][]int) []int {
	buf := make([]int, len(edges))
	for i := 0; i < len(edges); i++ {
		buf[i] = math.MaxInt
	}
	buf[start] = 0
	buf = bfs(buf, start, 0, edges)
	for i := 0; i < len(buf); i++ {
		if buf[i] == math.MaxInt {
			buf[i] = -1
		}
	}
	return buf
}

func bfs(dst []int, s, v int, e [][][]int) []int {
	ce := e[s]
	for i := 0; i < len(ce); i++ {
		c := ce[i]
		dst[c[0]] = min(dst[c[0]], v+c[1])
	}
	for i := 0; i < len(ce); i++ {
		c := ce[i]
		dst = bfs(dst, c[0], dst[c[0]], e)
	}
	return dst
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

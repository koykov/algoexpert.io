package count_squares

import "math"

func CountSquares(p [][]int) int {
	type adjd struct {
		idx  int
		dist float64
	}

	n := len(p)
	if n < 4 {
		return 0
	}
	adj := make(map[int][]adjd, len(p))
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			adj[i] = append(adj[i], adjd{idx: j, dist: dist(p[i], p[j])})
		}
	}
	return -1
}

func dist(a, b []int) float64 {
	return math.Sqrt(math.Pow(float64(b[0]-a[0]), 2) + math.Pow(float64(b[1]-a[1]), 2))
}

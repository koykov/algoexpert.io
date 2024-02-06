package knight_connection

import "math"

func KnightConnection(a []int, b []int) int {
	var x, y int
	if x, y = abs(a[0]-b[0]), abs(a[1]-b[1]); x+y == 1 {
		return 2
	}
	v := float64(move(x, y))
	return int(math.Ceil(v / 2))
}

func move(x, y int) int {
	if x <= 0 && y <= 0 {
		return 0
	}
	return move(max(x, y)-2, min(x, y)-1) + 1
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

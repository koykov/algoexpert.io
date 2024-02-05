package knight_connection

import "math"

func KnightConnection(a []int, b []int) int {
	v := walk(-8, 8, a[0], a[1], b[0], b[1])
	return int(math.Ceil(v / 2))
}

func walk(px, py, cx, cy, tx, ty int) float64 {
	if cx < -8 || cx > 8 || cy < -8 || cy > 8 || (cx == px && cy == py) {
		return 1000 // infinity
	}
	switch {
	case cx == tx && cy == ty:
		return 0
	case cx-2 == tx && cy+1 == ty,
		cx-1 == tx && cy+2 == ty,
		cx+1 == tx && cy+2 == ty,
		cx+2 == tx && cy+1 == ty,
		cx+2 == tx && cy-1 == ty,
		cx+1 == tx && cy-2 == ty,
		cx-1 == tx && cy-2 == ty,
		cx-2 == tx && cy-1 == ty:
		return 1
	}
	min := float64(1000) // infinity
	min = math.Min(min, walk(cx, cy, cx-2, cy+1, tx, ty))
	min = math.Min(min, walk(cx, cy, cx-1, cy+2, tx, ty))
	min = math.Min(min, walk(cx, cy, cx+1, cy+2, tx, ty))
	min = math.Min(min, walk(cx, cy, cx+2, cy+1, tx, ty))
	min = math.Min(min, walk(cx, cy, cx+2, cy-1, tx, ty))
	min = math.Min(min, walk(cx, cy, cx+1, cy-2, tx, ty))
	min = math.Min(min, walk(cx, cy, cx-1, cy-2, tx, ty))
	min = math.Min(min, walk(cx, cy, cx-2, cy-1, tx, ty))
	return min + 1
}

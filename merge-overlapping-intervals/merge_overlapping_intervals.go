package merge_overlapping_intervals

import (
	"math"
	"sort"
)

func MergeOverlappingIntervals(m [][]int) (o [][]int) {
	sort.Slice(m, func(i, j int) bool {
		return m[i][0] < m[j][0]
	})
	for i := 0; i < len(m); i++ {
		l, r := m[i][0], m[i][1]
		if l == -math.MaxInt {
			continue
		}
		for j := i + 1; j < len(m); j++ {
			l1, r1 := m[j][0], m[j][1]
			if l1 == -math.MaxInt {
				continue
			}
			if (l1 >= l && l1 <= r) || (r1 >= l && r1 <= r) {
				l = min(l, l1)
				r = max(r, r1)
				m[j][0] = -math.MaxInt
			}
		}
		o = append(o, []int{l, r})
	}
	return
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

package minimum_waiting_time

import "sort"

func MinimumWaitingTime(queries []int) (t int) {
	sort.Ints(queries)
	var min int
	for i := 1; i < len(queries); i++ {
		min += queries[i-1]
		t += min
	}
	return
}

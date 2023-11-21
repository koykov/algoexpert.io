package task_assignment

import "sort"

func TaskAssignment(_ int, tasks []int) (r [][]int) {
	type iv struct {
		i, v int
	}
	n := len(tasks)
	buf := make([]iv, 0, n)
	for i := 0; i < n; i++ {
		buf = append(buf, iv{i, tasks[i]})
	}
	sort.Slice(buf, func(i, j int) bool {
		return buf[i].v < buf[j].v
	})

	for i := 0; i < n/2; i++ {
		r = append(r, []int{buf[i].i, buf[n-i-1].i})
	}
	return
}

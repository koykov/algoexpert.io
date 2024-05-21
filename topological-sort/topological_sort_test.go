package topological_sort

import "testing"

func TestTopologicalSort(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		jobs_ := jobs(4)
		deps := []Dep{{1, 2}, {1, 3}, {3, 2}, {4, 2}, {4, 3}}
		order := TopologicalSort(jobs_, deps)
		if !isValidTopologicalOrder(order, jobs_, deps) {
			t.Fail()
		}
	})
	t.Run("1", func(t *testing.T) {
		//
	})
}

func jobs(n int) []int {
	var out []int
	for i := 1; i <= n; i++ {
		out = append(out, i)
	}
	return out
}

func isValidTopologicalOrder(order []int, jobs []int, deps []Dep) bool {
	visited := map[int]bool{}
	for _, candidate := range order {
		for _, dep := range deps {
			if _, found := visited[dep.Job]; found && candidate == dep.Prereq {
				return false
			}
		}
		visited[candidate] = true
	}
	for _, job := range jobs {
		if _, found := visited[job]; !found {
			return false
		}
	}
	return len(order) == len(jobs)
}

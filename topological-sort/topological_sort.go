package topological_sort

type Dep struct {
	Prereq int
	Job    int
}

func TopologicalSort(jobs []int, deps []Dep) []int {
	sol := solution{
		jobs:  jobs,
		deps:  deps,
		visit: make(map[int]bool),
		idx:   make(map[int]int),
	}
	for i := 0; i < len(jobs); i++ {
		if !sol.visit[jobs[i]] {
			sol.dfs(jobs[i])
		}
	}
	return sol.reverse()
}

type solution struct {
	jobs, stack []int
	deps        []Dep
	visit       map[int]bool
	idx         map[int]int
	cycle       bool
}

func (s *solution) dfs(v int) {
	if s.visit[v] {
		return
	}
	s.visit[v] = true
	for i := 0; i < len(s.deps); i++ {
		d := &s.deps[i]
		if d.Prereq == v {
			s.dfs(d.Job)
		}
	}
	s.stack = append(s.stack, v)
}

func (s *solution) reverse() []int {
	n := len(s.stack)
	for i := 0; i < n/2; i++ {
		s.stack[i], s.stack[n-i-1] = s.stack[n-i-1], s.stack[i]
	}
	return s.stack
}

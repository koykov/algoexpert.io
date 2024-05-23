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
	}
	for i := 0; i < len(jobs); i++ {
		if !sol.visit[jobs[i]] {
			sol.dfs(jobs[i])
		}
	}
	sol.reverse()
	if sol.valid() {
		return sol.stack
	}
	return []int{}
}

type solution struct {
	jobs, stack []int
	deps        []Dep
	visit       map[int]bool
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

func (s *solution) valid() bool {
	for k := range s.visit {
		delete(s.visit, k)
	}
	for _, candidate := range s.stack {
		for _, dep := range s.deps {
			if _, found := s.visit[dep.Job]; found && candidate == dep.Prereq {
				return false
			}
		}
		s.visit[candidate] = true
	}
	for _, job := range s.jobs {
		if _, found := s.visit[job]; !found {
			return false
		}
	}
	return len(s.stack) == len(s.jobs)
}

func (s *solution) reverse() []int {
	n := len(s.stack)
	for i := 0; i < n/2; i++ {
		s.stack[i], s.stack[n-i-1] = s.stack[n-i-1], s.stack[i]
	}
	return s.stack
}

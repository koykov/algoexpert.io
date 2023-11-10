package union_find

type UnionFind struct {
	p map[int]*int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{p: make(map[int]*int)}
}

func (dsu *UnionFind) CreateSet(v int) {
	dsu.p[v] = &v
}

func (dsu *UnionFind) Find(v int) *int {
	return dsu.p[v]
}

func (dsu *UnionFind) Union(a, b int) {
	ai, bi := dsu.Find(a), dsu.Find(b)
	if ai == nil || bi == nil {
		return
	}
	for k, v := range dsu.p {
		if dsu.p[b] == v && k != b {
			dsu.p[k] = dsu.p[a]
		}
	}
	dsu.p[b] = dsu.p[a]
}

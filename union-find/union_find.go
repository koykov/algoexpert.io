package union_find

import "math/rand"

const maxN = 1e6

type UnionFind struct {
	p []int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{p: make([]int, maxN)}
}

func (dsu *UnionFind) CreateSet(v int) {
	dsu.p[v] = v
}

func (dsu *UnionFind) Find(v int) *int {
	if dsu.p[v] == v {
		return &dsu.p[v]
	}
	dsu.p[v] = *dsu.Find(dsu.p[v])
	return &dsu.p[v]
}

func (dsu *UnionFind) Union(a, b int) {
	ai, bi := dsu.Find(a), dsu.Find(b)
	if rand.Int()%2 == 0 {
		ai, bi = bi, ai
	}
	dsu.p[*ai] = *bi
}

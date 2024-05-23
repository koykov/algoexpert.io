package kruskals_algorithm

import "sort"

func KruskalsAlgorithm(edges [][][]int) [][][]int {
	var cpy [][]int
	for i := 0; i < len(edges); i++ {
		v := edges[i]
		for j := 0; j < len(v); j++ {
			e := v[j]
			if e[0] > i {
				cpy = append(cpy, []int{i, e[0], e[1]})
			}
		}
	}
	sort.Slice(cpy, func(i, j int) bool {
		return cpy[i][2] < cpy[j][2]
	})

	var buf [][][]int
	union := newDSU(len(edges))

	for i := 0; i < len(edges); i++ {
		union.add(i)
		buf = append(buf, [][]int{})
	}

	for i := 0; i < len(cpy); i++ {
		e := cpy[i]
		root0, root1 := union.find(e[0]), union.find(e[1])
		if root0 != root1 {
			buf[e[0]] = append(buf[e[0]], []int{e[1], e[2]})
			buf[e[1]] = append(buf[e[1]], []int{e[0], e[2]})
		}
	}
	return buf
}

// disjoint set union
type dsu struct {
	p, r []int
}

func newDSU(len int) *dsu {
	return &dsu{make([]int, len), make([]int, len)}
}

func (u *dsu) add(val int) {
	u.p[val], u.r[val] = val, 0
}

func (u *dsu) find(val int) int {
	if val != u.p[val] {
		u.p[val] = u.find(u.p[val])
	}
	return u.p[val]
}

func (u *dsu) union(val0, val1 int) {
	root0, root1 := u.find(val0), u.find(val1)
	if u.r[root0] < u.r[root1] {
		u.p[root0] = root1
	} else if u.r[root0] < u.r[root1] {
		u.p[root1] = root0
	} else {
		u.p[root1] = root0
		u.r[root0]++
	}
}

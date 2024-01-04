package group_anagrams

import "sort"

func GroupAnagrams(w []string) [][]string {
	var buf []byte
	var bs_ bs
	reg := make(map[string][]string, len(w))
	for i := 0; i < len(w); i++ {
		buf = append(buf[:0], w[i]...)
		bs_.p = buf
		sort.Sort(&bs_)
		k := string(buf)
		reg[k] = append(reg[k], w[i])
	}
	r := make([][]string, 0, len(reg))
	for _, set := range reg {
		r = append(r, set)
	}
	return r
}

type bs struct {
	p []byte
}

func (b *bs) Len() int {
	return len(b.p)
}

func (b *bs) Less(i, j int) bool {
	return b.p[i] < b.p[j]
}

func (b *bs) Swap(i, j int) {
	b.p[i], b.p[j] = b.p[j], b.p[i]
}

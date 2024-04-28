package disk_stacking

import "sort"

func DiskStacking(d [][]int) [][]int {
	sort.Slice(d, func(i, j int) bool {
		return d[i][2] < d[j][2]
	})

	h := make([]int, len(d))
	seq := make([]int, len(d))
	for i := 0; i < len(d); i++ {
		h[i] = d[i][2]
		seq[i] = -1
	}

	var hm, hmi int
	for i := 0; i < len(d); i++ {
		for j := 0; j < i; j++ {
			c, o := d[i], d[j]
			if o[0] < c[0] && o[1] < c[1] && o[2] < c[2] {
				if t := c[2] + h[j]; h[i] < t {

					h[i] = t
					seq[i] = j
				}
			}
		}
		if h[i] > hm {
			hm, hmi = h[i], i
		}
	}

	_ = hmi
	r := make([][]int, 0, len(d))
	for hmi != -1 {
		r = append(r, d[hmi])
		hmi = seq[hmi]
	}
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}
	return r
}

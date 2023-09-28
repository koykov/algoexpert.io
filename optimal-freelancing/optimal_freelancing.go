package optimal_freelancing

import "sort"

func OptimalFreelancing(jobs []map[string]int) int {
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i]["payment"] > jobs[j]["payment"]
	})
	var d [7]int
	for i := 0; i < len(jobs); i++ {
		job := jobs[i]
		dl := job["deadline"]
		if dl > 7 {
			dl = 7
		}
		if d[dl-1] == 0 {
			d[dl-1] = job["payment"]
		} else {
			for j := dl; j > 0; j-- {
				if p := job["payment"]; d[j-1] < p {
					d[j-1] = p
				}
			}
		}
	}
	return d[0] + d[1] + d[2] + d[3] + d[4] + d[5] + d[6]
}

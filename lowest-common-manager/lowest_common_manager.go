package lowest_common_manager

type OrgChart struct {
	Name          string
	DirectReports []*OrgChart
}

func GetLowestCommonManager(root, r0, r1 *OrgChart) *OrgChart {
	var s solution
	s.dfs(root, r0, r1)
	return s.r
}

type solution struct {
	r      *OrgChart
	a0, a1 bool
}

func (s *solution) dfs(root *OrgChart, r0, r1 *OrgChart) {
	if root.Name == r0.Name {
		s.a0 = true
		return
	}
	if root.Name == r1.Name {
		s.a1 = true
		return
	}
	for i := 0; i < len(root.DirectReports); i++ {
		s.dfs(root.DirectReports[i], r0, r1)
		if s.a0 && s.a1 && s.r == nil {
			s.r = root
			break
		}
	}
}

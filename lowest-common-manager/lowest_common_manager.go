package lowest_common_manager

type OrgChart struct {
	Name          string
	DirectReports []*OrgChart
}

func GetLowestCommonManager(root, r0, r1 *OrgChart) *OrgChart {
	var x *OrgChart
	var c int
	if root == nil || r0.Name == root.Name || r1.Name == root.Name {
		return root
	}
	for i := 0; i < len(root.DirectReports); i++ {
		r := GetLowestCommonManager(root.DirectReports[i], r0, r1)
		if r != nil {
			c++
			x = r
		}
	}
	if c == 2 {
		return root
	}
	return x
}

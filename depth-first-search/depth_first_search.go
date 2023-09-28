package depth_first_search

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	return dive(array, n)
}

func dive(dst []string, root *Node) []string {
	dst = append(dst, root.Name)
	for i := 0; i < len(root.Children); i++ {
		dst = dive(dst, root.Children[i])
	}
	return dst
}

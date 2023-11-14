package breadth_first_search

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) BreadthFirstSearch(dst []string) []string {
	var buf []*Node
	buf = append(buf, n)
	for len(buf) > 0 {
		c := buf[0]
		buf = append(buf[:0], buf[1:]...)
		dst = append(dst, c.Name)
		for i := 0; i < len(c.Children); i++ {
			buf = append(buf, c.Children[i])
		}
	}
	return dst
}

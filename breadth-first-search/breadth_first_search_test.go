package breadth_first_search

import "testing"

func TestBreadthFirstSearch(t *testing.T) {
	root := &Node{
		Name: "A",
		Children: []*Node{
			{
				Name: "B",
				Children: []*Node{
					{Name: "E"},
					{Name: "F", Children: []*Node{
						{Name: "I"},
						{Name: "J"},
					}},
				},
			},
			{Name: "C"},
			{Name: "D", Children: []*Node{
				{Name: "G", Children: []*Node{
					{Name: "K"},
				}},
				{Name: "H"},
			}},
		},
	}
	r := root.BreadthFirstSearch(nil)
	t.Log(r)
}

package number_of_ways_to_traverse_graph

func NumberOfWaysToTraverseGraph(w, h int) int {
	if w == 1 || h == 1 {
		return 1
	}
	return NumberOfWaysToTraverseGraph(w-1, h) + NumberOfWaysToTraverseGraph(w, h-1)
}

package number_of_ways_to_make_change

func NumberOfWaysToMakeChange(n int, denoms []int) int {
	switch {
	case n == 0:
		return 1
	case n < 0 || len(denoms) == 0:
		return 0
	default:
		d := denoms[0]
		return NumberOfWaysToMakeChange(n-d, denoms) + NumberOfWaysToMakeChange(n, denoms[1:])
	}
}

package non_constructible_change

import "sort"

func NonConstructibleChange(coins []int) int {
	if len(coins) == 0 {
		return 1
	}
	sort.Ints(coins)
	c := 0
	for i := 0; i < len(coins); i++ {
		if coins[i] <= c+1 {
			c += coins[i]
			continue
		}
		break
	}
	return c + 1
}

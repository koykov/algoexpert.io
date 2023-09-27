package max_profit_with_k_transactions

import "math"

func MaxProfitWithKTransactions(prices []int, k int) int {
	n := len(prices)
	if n <= 0 {
		return 0
	}
	buf := make([]int, (n+1)*(k+1))
	profit := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		profit[i] = buf[i*(n+1) : i*(n+1)+n+1]
	}

	for i := 1; i <= k; i++ {
		p := -math.MaxInt
		for j := 1; j < n; j++ {
			p = max(p, profit[i-1][j-1]-prices[j-1])
			profit[i][j] = max(profit[i][j-1], prices[j]+p)
		}
	}
	return profit[k][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

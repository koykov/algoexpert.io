package max_profit_with_k_transactions

import (
	"strconv"
	"testing"
)

var stages = []struct {
	prices []int
	k, r   int
}{
	{
		prices: []int{5, 11, 3, 50, 60, 90},
		k:      2,
		r:      93,
	},
	{
		prices: []int{},
		k:      1,
		r:      0,
	},
	{
		prices: []int{1},
		k:      1,
		r:      0,
	},
	{
		prices: []int{1, 10},
		k:      1,
		r:      9,
	},
	{
		prices: []int{1, 10},
		k:      3,
		r:      9,
	},
	{
		prices: []int{3, 2, 5, 7, 1, 3, 7},
		k:      1,
		r:      6,
	},
	{
		prices: []int{5, 11, 3, 50, 60, 90},
		k:      3,
		r:      93,
	},
	{
		prices: []int{5, 11, 3, 50, 40, 90},
		k:      2,
		r:      97,
	},
	{
		prices: []int{5, 11, 3, 50, 40, 90},
		k:      3,
		r:      103,
	},
	{
		prices: []int{50, 25, 12, 4, 3, 10, 1, 100},
		k:      2,
		r:      106,
	},
	{
		prices: []int{100, 99, 98, 97, 1},
		k:      5,
		r:      0,
	},
	{
		prices: []int{1, 100, 2, 200, 3, 300, 4, 400, 5, 500},
		k:      5,
		r:      1485,
	},
	{
		prices: []int{1, 100, 101, 200, 201, 300, 301, 400, 401, 500},
		k:      5,
		r:      499,
	},
	{
		prices: []int{1, 25, 24, 23, 12, 36, 14, 40, 31, 41, 5},
		k:      4,
		r:      84,
	},
	{
		prices: []int{1, 25, 24, 23, 12, 36, 14, 40, 31, 41, 5},
		k:      2,
		r:      62,
	},
}

func TestMaxProfitWithKTransactions(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := MaxProfitWithKTransactions(stage.prices, stage.k)
			if r != stage.r {
				t.FailNow()
			}
		})
	}
}

func BenchmarkMaxProfitWithKTransactions(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stage := &stages[i%len(stages)]
		MaxProfitWithKTransactions(stage.prices, stage.k)
	}
}

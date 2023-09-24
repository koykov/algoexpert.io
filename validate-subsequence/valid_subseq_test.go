package validate_subsequence

import (
	"fmt"
	"testing"
)

var stages = []struct {
	array, sequence []int
	valid           bool
}{
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{1, 6, -1, 10},
		true,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		true,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{5, 1, 22, 6, -1, 8, 10},
		true,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{22, 25, 6},
		true,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{1, 6, 10},
		true,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{5, 1, 22, 10},
		true,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{5, -1, 8, 10},
		true,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{25},
		true,
	},
	{
		[]int{1, 1, 1, 1, 1},
		[]int{1, 1, 1},
		true,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{5, 1, 22, 25, 6, -1, 8, 10, 12},
		false,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{4, 5, 1, 22, 25, 6, -1, 8, 10},
		false,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{5, 1, 22, 23, 6, -1, 8, 10},
		false,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{5, 1, 22, 22, 25, 6, -1, 8, 10},
		false,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{5, 1, 22, 22, 6, -1, 8, 10},
		false,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{1, 6, -1, -1},
		false,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{1, 6, -1, -1, 10},
		false,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{1, 6, -1, -2},
		false,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{26},
		false,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{5, 1, 25, 22, 6, -1, 8, 10},
		false,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{5, 26, 22, 8},
		false,
	},
	{
		[]int{1, 1, 6, 1},
		[]int{1, 1, 1, 6},
		false,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{1, 6, -1, 10, 11, 11, 1, 11},
		false,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{5, 1, 22, 25, 6, -1, 8, 10, 10},
		false,
	},
	{
		[]int{5, 1, 22, 25, 6, -1, 8, 10},
		[]int{1, 6, -1, 5},
		false,
	},
}

func TestIsValidSubsequence(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run(fmt.Sprintf("trivial/%d", i), func(t *testing.T) {
			v := IsValidSubsequence(stage.array, stage.sequence)
			if v != stage.valid {
				t.FailNow()
			}
		})
	}
}

func BenchmarkIsValidSubsequence(b *testing.B) {
	b.Run("trivial", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			stage := &stages[i%len(stages)]
			IsValidSubsequence(stage.array, stage.sequence)
		}
	})
}

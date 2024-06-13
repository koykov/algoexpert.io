package largest_rectangle_under_skyline

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLargestRectangleUnderSkyline(t *testing.T) {
	stages := []struct {
		b []int
		r int
	}{
		{b: []int{1, 3, 3, 2, 4, 1, 5, 3, 2}, r: 9},
		{b: []int{4, 4, 4, 2, 2, 1}, r: 12},
		{b: []int{1, 3, 3, 2, 4, 1, 5, 3}, r: 8},
		{b: []int{5, 5, 2, 2, 4, 1}, r: 10},
		{b: []int{1, 2, 3, 4, 5, 11}, r: 12},
		{b: []int{25, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, r: 25},
		{b: []int{20, 2, 2, 2, 2, 2, 10, 5, 5, 5, 4, 4}, r: 24},
		{b: []int{5, 10, 5, 15, 10, 25}, r: 30},
		{b: []int{1, 1, 1, 1}, r: 4},
		{b: []int{10, 21}, r: 21},
		{b: []int{11, 21}, r: 22},
		{b: []int{3, 3, 3, 4, 4, 4, 1, 3, 1, 2, 8, 9, 1}, r: 18},
		{b: []int{5}, r: 5},
		{b: []int{10, 1, 2, 3, 4, 5, 6, 7}, r: 16},
		{b: []int{10, 1, 2, 3, 3, 5, 6, 7}, r: 15},
		{b: []int{}, r: 0},
	}
	for i, tt := range stages {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			input := tt.b
			expected := tt.r
			actual := LargestRectangleUnderSkyline(input)
			require.Equal(t, expected, actual)
		})
	}
}

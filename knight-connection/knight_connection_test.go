package knight_connection

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKnightConnection(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		knightA := []int{0, 0}
		knightB := []int{2, 1}
		expected := 1
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("1", func(t *testing.T) {
		knightA := []int{15, -12}
		knightB := []int{15, -12}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("2", func(t *testing.T) {
		knightA := []int{1, 0}
		knightB := []int{0, 0}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("3", func(t *testing.T) {
		knightA := []int{1, 0}
		knightB := []int{0, 0}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("4", func(t *testing.T) {
		knightA := []int{0, 0}
		knightB := []int{0, 1}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("5", func(t *testing.T) {
		knightA := []int{1, 1}
		knightB := []int{0, 0}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("6", func(t *testing.T) {
		knightA := []int{0, 0}
		knightB := []int{-1, -1}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("7", func(t *testing.T) {
		knightA := []int{2, 1}
		knightB := []int{0, 0}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("8", func(t *testing.T) {
		knightA := []int{3, 3}
		knightB := []int{0, 0}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("9", func(t *testing.T) {
		knightA := []int{2, 1}
		knightB := []int{-1, -2}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("10", func(t *testing.T) {
		knightA := []int{2, 1}
		knightB := []int{-2, -4}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("11", func(t *testing.T) {
		knightA := []int{5, 2}
		knightB := []int{-2, -4}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("12", func(t *testing.T) {
		knightA := []int{10, 10}
		knightB := []int{-10, -10}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("13", func(t *testing.T) {
		knightA := []int{15, 15}
		knightB := []int{-10, -10}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("14", func(t *testing.T) {
		knightA := []int{-15, 2}
		knightB := []int{-3, 20}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("15", func(t *testing.T) {
		knightA := []int{20, 20}
		knightB := []int{0, 0}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
	t.Run("16", func(t *testing.T) {
		knightA := []int{18, -13}
		knightB := []int{0, 12}
		expected := 0
		actual := KnightConnection(knightA, knightB)
		require.Equal(t, expected, actual)
	})
}

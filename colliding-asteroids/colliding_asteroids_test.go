package colliding_asteroids

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCollidingAsteroids(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		input := []int{-3, 5, -8, 6, 7, -4, -7}
		expected := []int{-3, -8, 6}
		actual := CollidingAsteroids(input)
		require.Equal(t, expected, actual)
	})
	t.Run("1", func(t *testing.T) {
		input := []int{1, 2, 3, -4}
		expected := []int{-4}
		actual := CollidingAsteroids(input)
		require.Equal(t, expected, actual)
	})
	t.Run("3", func(t *testing.T) {
		input := []int{-70, 100, 23, 42, -50, -75, 1, -2, -3}
		expected := []int{-70, 100}
		actual := CollidingAsteroids(input)
		require.Equal(t, expected, actual)
	})
}

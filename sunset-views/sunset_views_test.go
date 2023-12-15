package sunset_views

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSunsetViews(t *testing.T) {
	buildings := []int{3, 5, 4, 4, 3, 1, 3, 2}
	direction := "EAST"
	expected := []int{1, 3, 6, 7}
	actual := SunsetViews(buildings, direction)
	require.Equal(t, expected, actual)
}

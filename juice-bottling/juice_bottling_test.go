package juice_bottling

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJuiceBottling(t *testing.T) {
	input := []int{0, 2, 5, 6}
	expected := []int{1, 2}
	actual := JuiceBottling(input)
	require.Equal(t, expected, actual)
}

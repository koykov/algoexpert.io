package min_rewards

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinRewards(t *testing.T) {
	input := []int{8, 4, 2, 1, 3, 6, 7, 9, 5}
	output := MinRewards(input)
	expected := 25
	require.Equal(t, expected, output)
}

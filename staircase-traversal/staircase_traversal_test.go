package staircase_traversal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStaircaseTraversal(t *testing.T) {
	stairs := 4
	maxSteps := 2
	expected := 5
	actual := StaircaseTraversal(stairs, maxSteps)
	require.Equal(t, expected, actual)
}

package blackjack_probability

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlackjackProbability(t *testing.T) {
	target := 21
	startingHand := 15
	expected := 0.45
	actual := BlackjackProbability(target, startingHand)
	require.Equal(t, expected, actual)
}

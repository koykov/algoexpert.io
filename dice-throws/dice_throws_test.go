package dice_throws

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiceThrows(t *testing.T) {
	numDice := 2
	numSides := 6
	target := 7
	expected := 6
	actual := DiceThrows(numDice, numSides, target)
	require.Equal(t, expected, actual)
}

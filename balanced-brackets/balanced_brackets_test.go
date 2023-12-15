package balanced_brackets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBalancedBrackets(t *testing.T) {
	expected := true
	output := BalancedBrackets("([])(){}(())()()")
	// output := BalancedBrackets("(()agwg())((()agwga()())gawgwgag)")
	require.Equal(t, expected, output)
}

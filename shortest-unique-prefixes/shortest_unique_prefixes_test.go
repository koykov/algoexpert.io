package shortest_unique_prefixes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShortestUniquePrefixes(t *testing.T) {
	strings := []string{"algoexpert", "algorithm", "frontendexpert", "mlexpert"}
	expected := []string{"algoe", "algor", "f", "m"}
	actual := ShortestUniquePrefixes(strings)
	require.Equal(t, expected, actual)
}

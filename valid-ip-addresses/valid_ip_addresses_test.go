package valid_ip_addresses

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidIPAddresses(t *testing.T) {
	t.Run("1921680", func(t *testing.T) {
		input := "1921680"
		expected := []string{
			"1.9.216.80",
			"1.92.16.80",
			"1.92.168.0",
			"19.2.16.80",
			"19.2.168.0",
			"19.21.6.80",
			"19.21.68.0",
			"19.216.8.0",
			"192.1.6.80",
			"192.1.68.0",
			"192.16.8.0",
		}
		actual := ValidIPAddresses(input)
		require.ElementsMatch(t, expected, actual)
	})
	t.Run("3700100", func(t *testing.T) {
		input := "3700100"
		expected := []string{
			"3.70.0.100",
			"37.0.0.100",
		}
		actual := ValidIPAddresses(input)
		require.ElementsMatch(t, expected, actual)
	})
}

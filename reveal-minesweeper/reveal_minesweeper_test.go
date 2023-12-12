package reveal_minesweeper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRevealMinesweeper(t *testing.T) {
	board := [][]string{
		{"H", "H", "H", "H", "M"},
		{"H", "1", "M", "H", "1"},
		{"H", "H", "H", "H", "H"},
		{"H", "H", "H", "H", "H"},
	}
	row := 3
	column := 4
	expected := [][]string{
		{"0", "1", "H", "H", "M"},
		{"0", "1", "M", "2", "1"},
		{"0", "1", "1", "1", "0"},
		{"0", "0", "0", "0", "0"},
	}
	actual := RevealMinesweeper(board, row, column)
	require.Equal(t, expected, actual)
}

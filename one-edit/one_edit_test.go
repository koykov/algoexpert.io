package one_edit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOneEdit(t *testing.T) {
	stringOne := "hello"
	stringTwo := "helo"
	actual := OneEdit(stringOne, stringTwo)
	require.Equal(t, true, actual)
}

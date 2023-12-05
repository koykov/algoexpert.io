package merging_linked_lists

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergingLinkedLists(t *testing.T) {
	l1 := &LinkedList{Value: 1}
	l1.Next = &LinkedList{Value: 2}
	l2 := &LinkedList{Value: 3}
	l2.Next = l1.Next

	expected := l1.Next
	actual := MergingLinkedLists(l1, l2)
	require.Equal(t, expected, actual)
}

package merge_linked_lists

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func NewLinkedList(val int, others ...int) *LinkedList {
	ll := &LinkedList{Value: val}
	current := ll
	for _, other := range others {
		current.Next = &LinkedList{Value: other}
		current = current.Next
	}
	return ll
}

func (ll *LinkedList) ToArray() []int {
	vals := []int{}
	current := ll
	for current != nil {
		vals = append(vals, current.Value)
		current = current.Next
	}
	return vals
}

func TestMergeLinkedLists(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		list1 := NewLinkedList(2, 6, 7, 8)
		list2 := NewLinkedList(1, 3, 4, 5, 9, 10)
		output := MergeLinkedLists(list1, list2)
		expectedNodes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		require.Equal(t, expectedNodes, output.ToArray())
	})
	t.Run("1", func(t *testing.T) {
		list1 := NewLinkedList(1, 1, 1, 3, 4, 5, 5, 5, 5, 10)
		list2 := NewLinkedList(1, 1, 2, 2, 5, 6, 10, 10)
		output := MergeLinkedLists(list1, list2)
		expectedNodes := []int{1, 1, 1, 1, 1, 2, 2, 3, 4, 5, 5, 5, 5, 5, 6, 10, 10, 10}
		require.Equal(t, expectedNodes, output.ToArray())
	})
	t.Run("2", func(t *testing.T) {
		list1 := NewLinkedList(2)
		list2 := NewLinkedList(1)
		output := MergeLinkedLists(list1, list2)
		expectedNodes := []int{1, 2}
		require.Equal(t, expectedNodes, output.ToArray())
	})
}

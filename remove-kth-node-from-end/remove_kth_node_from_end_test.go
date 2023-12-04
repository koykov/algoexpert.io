package remove_kth_node_from_end

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func NewLinkedList(root int, children ...int) *LinkedList {
	ll := &LinkedList{root, nil}
	ll.Add(children...)
	return ll
}

func (ll *LinkedList) Add(values ...int) {
	current := ll
	for current.Next != nil {
		current = current.Next
	}
	for _, value := range values {
		current.Next = &LinkedList{value, nil}
		current = current.Next
	}
}

func (ll *LinkedList) ToArray() []int {
	output, current := []int{ll.Value}, ll
	for current.Next != nil {
		current = current.Next
		output = append(output, current.Value)
	}
	return output
}

func TestRemoveKthNodeFromEnd(t *testing.T) {
	ll := NewLinkedList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	RemoveKthNodeFromEnd(ll, 10)
	output, expected := ll.ToArray(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	require.Equal(t, expected, output)
}

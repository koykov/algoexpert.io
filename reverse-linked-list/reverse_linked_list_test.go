package reverse_linked_list

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

func TestReverseLinkedList(t *testing.T) {
	input := NewLinkedList(0, 1, 2, 3, 4, 5)
	output := ReverseLinkedList(input)
	expected := []int{5, 4, 3, 2, 1, 0}
	require.NotNil(t, output)
	require.Equal(t, expected, output.ToArray())
}

package linked_list_construction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDoublyLinkedList(t *testing.T) {
	// nn := []*Node{
	// 	{Value: 1},
	// 	{Value: 2},
	// 	{Value: 3},
	// 	{Value: 4},
	// 	{Value: 5},
	// }
	// l := &DoublyLinkedList{
	// 	Head: nn[0],
	// 	Tail: nn[4],
	// }
	// l.Head.Next = nn[1]
	// l.Head.Next.Prev = nn[0]
	// l.Head.Next.Next = nn[2]
	// l.Head.Next.Next.Prev = nn[1]
	// l.Head.Next.Next.Next = nn[3]
	// l.Head.Next.Next.Next.Prev = nn[2]
	// l.Head.Next.Next.Next.Next = nn[4]
	// l.Head.Next.Next.Next.Next.Prev = nn[3]
	// t.Log(l.String())
	// l.SetHead(&Node{Value: 4})
	// t.Log(l.String())
	// l.SetTail(&Node{Value: 6})
	// t.Log(l.String())
	// l.InsertBefore(&Node{Value: 6}, &Node{Value: 3})
	// t.Log(l.String())
	// l.InsertAfter(&Node{Value: 6}, &Node{Value: 3})
	// t.Log(l.String())
	// // l.InsertAtPosition(1, &Node{Value: 3})
	// // t.Log(l.String())
	linkedList := NewDoublyLinkedList()
	one := NewNode(1)
	two := NewNode(2)
	three := NewNode(3)
	three2 := NewNode(3)
	three3 := NewNode(3)
	four := NewNode(4)
	five := NewNode(5)
	six := NewNode(6)
	bindNodes(one, two)
	bindNodes(two, three)
	bindNodes(three, four)
	bindNodes(four, five)
	linkedList.Head = one
	linkedList.Tail = five

	linkedList.SetHead(four)
	require.Equal(t, getNodeValuesHeadToTail(linkedList), []int{4, 1, 2, 3, 5})
	require.Equal(t, getNodeValuesTailToHead(linkedList), []int{5, 3, 2, 1, 4})

	linkedList.SetTail(six)
	require.Equal(t, getNodeValuesHeadToTail(linkedList), []int{4, 1, 2, 3, 5, 6})
	require.Equal(t, getNodeValuesTailToHead(linkedList), []int{6, 5, 3, 2, 1, 4})

	linkedList.InsertBefore(six, three)
	require.Equal(t, getNodeValuesHeadToTail(linkedList), []int{4, 1, 2, 5, 3, 6})
	require.Equal(t, getNodeValuesTailToHead(linkedList), []int{6, 3, 5, 2, 1, 4})

	linkedList.InsertAfter(six, three2)
	require.Equal(t, getNodeValuesHeadToTail(linkedList), []int{4, 1, 2, 5, 3, 6, 3})
	require.Equal(t, getNodeValuesTailToHead(linkedList), []int{3, 6, 3, 5, 2, 1, 4})

	linkedList.InsertAtPosition(1, three3)
	require.Equal(t, getNodeValuesHeadToTail(linkedList), []int{3, 4, 1, 2, 5, 3, 6, 3})
	require.Equal(t, getNodeValuesTailToHead(linkedList), []int{3, 6, 3, 5, 2, 1, 4, 3})

	linkedList.RemoveNodesWithValue(3)
	require.Equal(t, getNodeValuesHeadToTail(linkedList), []int{4, 1, 2, 5, 6})
	require.Equal(t, getNodeValuesTailToHead(linkedList), []int{6, 5, 2, 1, 4})

	linkedList.Remove(two)
	require.Equal(t, getNodeValuesHeadToTail(linkedList), []int{4, 1, 5, 6})
	require.Equal(t, getNodeValuesTailToHead(linkedList), []int{6, 5, 1, 4})

	require.Equal(t, linkedList.ContainsNodeWithValue(5), true)
}

func NewNode(value int) *Node { return &Node{Value: value} }

func getNodeValuesHeadToTail(ll *DoublyLinkedList) []int {
	values := []int{}
	node := ll.Head
	for node != nil {
		values = append(values, node.Value)
		node = node.Next
	}
	return values
}

func getNodeValuesTailToHead(ll *DoublyLinkedList) []int {
	values := []int{}
	node := ll.Tail
	for node != nil {
		values = append(values, node.Value)
		node = node.Prev
	}
	return values
}

func bindNodes(nodeOne *Node, nodeTwo *Node) {
	nodeOne.Next = nodeTwo
	nodeTwo.Prev = nodeOne
}

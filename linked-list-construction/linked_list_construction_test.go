package linked_list_construction

import "testing"

func TestNewDoublyLinkedList(t *testing.T) {
	nn := []*Node{
		{Value: 1},
		{Value: 2},
		{Value: 3},
		{Value: 4},
		{Value: 5},
	}
	l := &DoublyLinkedList{
		Head: nn[0],
		Tail: nn[4],
	}
	l.Head.Next = nn[1]
	l.Head.Next.Prev = nn[0]
	l.Head.Next.Next = nn[2]
	l.Head.Next.Next.Prev = nn[1]
	l.Head.Next.Next.Next = nn[3]
	l.Head.Next.Next.Next.Prev = nn[2]
	l.Head.Next.Next.Next.Next = nn[4]
	l.Head.Next.Next.Next.Next.Prev = nn[3]
	t.Log(l.String())
	l.SetHead(&Node{Value: 4})
	t.Log(l.String())
	l.SetTail(&Node{Value: 6})
	t.Log(l.String())
	l.InsertBefore(&Node{Value: 6}, &Node{Value: 3})
	t.Log(l.String())
	l.InsertAfter(&Node{Value: 6}, &Node{Value: 3})
	t.Log(l.String())
	l.InsertAtPosition(1, &Node{Value: 3})
	t.Log(l.String())
}

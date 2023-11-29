package linked_list_construction

import "strconv"

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (l *DoublyLinkedList) SetHead(node *Node) {
	pn := l.Head
	for {
		if pn = pn.Next; pn == nil || pn.Value == node.Value {
			break
		}
	}
	if pn == nil {
		node.Next = l.Head
		l.Head = node
		return
	}
	oh := l.Head
	l.Head = pn
	t := pn.Prev
	pn.Prev.Next = pn.Next
	pn.Next.Prev = t
	l.Head.Prev = nil
	l.Head.Next = oh
}

func (l *DoublyLinkedList) SetTail(node *Node) {
	pn := l.Head
	for {
		if pn = pn.Next; pn == nil || pn.Value == node.Value {
			break
		}
	}
	if pn == nil {
		node.Prev = l.Tail
		l.Tail.Next = node
		l.Tail = node
		return
	}
	ot := l.Tail
	l.Tail = pn
	t := pn.Prev
	pn.Prev.Next = pn.Next
	pn.Next.Prev = t
	l.Tail.Prev = nil
	l.Tail.Next = ot
}

func (l *DoublyLinkedList) InsertBefore(node, node1 *Node) {
	pn := l.Head
	for {
		if pn = pn.Next; pn == nil || pn.Value == node.Value {
			break
		}
	}
	if pn == nil {
		return
	}
	node1.Prev, node1.Next = pn.Prev, pn
	pn.Prev.Next = node1
	pn.Prev = node1
}

func (l *DoublyLinkedList) InsertAfter(node, node1 *Node) {
	pn := l.Head
	for {
		if pn = pn.Next; pn == nil || pn.Value == node.Value {
			break
		}
	}
	if pn == nil {
		return
	}
	node1.Prev, node1.Next = pn, pn.Next
	pn.Next = node1
}

func (l *DoublyLinkedList) InsertAtPosition(p int, node *Node) {
	pn := l.Head
	for i := 0; i < p; i++ {
		if pn = pn.Next; pn == nil {
			return
		}
	}
	node.Prev, node.Next = pn.Prev, pn
	pn.Prev.Next = node
}

func (l *DoublyLinkedList) RemoveNodesWithValue(value int) {
}

func (l *DoublyLinkedList) Remove(node *Node) {
}

func (l *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	return false
}

func (l *DoublyLinkedList) String() (r string) {
	r += strconv.Itoa(l.Head.Value)
	n := l.Head
	for {
		n = n.Next
		if n == nil {
			break
		}
		r += " <-> "
		r += strconv.Itoa(n.Value)
	}
	return
}

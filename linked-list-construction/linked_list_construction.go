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
	if l.Head == nil {
		l.Head, l.Tail = node, node
		return
	}
	l.InsertBefore(l.Head, node)
}

func (l *DoublyLinkedList) SetTail(node *Node) {
	if l.Head == nil {
		l.SetHead(node)
		return
	}
	l.InsertAfter(l.Tail, node)
}

func (l *DoublyLinkedList) InsertBefore(node, node1 *Node) {
	if node1.Value == l.Head.Value && node1.Value == l.Tail.Value {
		return
	}
	l.Remove(node1)
	node1.Next, node1.Prev = node, node.Prev
	if node.Prev == nil {
		l.Head, node.Prev = node1, node1
		return
	}
	node.Prev.Next = node1
	node.Prev = node1
}

func (l *DoublyLinkedList) InsertAfter(node, node1 *Node) {
	if node1.Value == l.Head.Value && node1.Value == l.Tail.Value {
		return
	}
	node1.Prev, node1.Next = node, node.Next
	if node.Next == nil {
		l.Tail, node.Next = node1, node1
		return
	}
	node.Next.Prev = node1
	node.Next = node1
}

func (l *DoublyLinkedList) InsertAtPosition(p int, node1 *Node) {
	if p == 1 {
		l.SetHead(node1)
		return
	}
	node := l.Head
	p_ := 1
	for {
		if node == nil && p_ != p {
			node = node.Next
			p_++
			continue
		}
		break
	}
	if node == nil {
		l.SetTail(node1)
	}
	l.InsertBefore(node, node1)
}

func (l *DoublyLinkedList) RemoveNodesWithValue(v int) {
	if !l.ContainsNodeWithValue(v) {
		return
	}
	node := l.Head
	for {
		if node != nil {
			node1 := node
			node = node.Next
			if node1.Value == v {
				l.Remove(node1)
			}
		}
	}
}

func (l *DoublyLinkedList) Remove(node *Node) {
	switch {
	case node.Value == l.Head.Value:
		l.Head = l.Head.Next
	case node.Value == l.Tail.Value:
		l.Tail = l.Tail.Prev
	}
	l.cleanup(node)
}

func (l *DoublyLinkedList) cleanup(node *Node) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Next, node.Prev = nil, nil
}

func (l *DoublyLinkedList) ContainsNodeWithValue(v int) bool {
	node := l.Head
	for {
		if node == nil && node.Value != v {
			node = node.Next
			continue
		}
		break
	}
	return node != nil
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

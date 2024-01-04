package linked_list_construction

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

func (ll *DoublyLinkedList) SetHead(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}

	ll.InsertBefore(ll.Head, node)
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
	if ll.Tail == nil {
		ll.Head = node
		ll.Tail = node
		return
	}

	ll.InsertAfter(ll.Tail, node)
}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	if nodeToInsert == node || nodeToInsert.Next == node {
		return
	}
	ll.Remove(nodeToInsert)

	if node == ll.Head {
		ll.Head = nodeToInsert
	} else {
		node.Prev.Next = nodeToInsert
	}

	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node
	node.Prev = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	if nodeToInsert == node || nodeToInsert.Prev == node {
		return
	}
	ll.Remove(nodeToInsert)

	if node == ll.Tail {
		ll.Tail = nodeToInsert
	} else {
		node.Next.Prev = nodeToInsert
	}

	nodeToInsert.Next = node.Next
	nodeToInsert.Prev = node
	node.Next = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	if position == 1 {
		ll.SetHead(nodeToInsert)
		return
	}

	current := ll.Head
	for i := 2; i < position; i++ {
		current = current.Next
	}
	ll.InsertAfter(current, nodeToInsert)
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	current := ll.Head

	for current != nil {
		tempNext := current.Next
		if current.Value == value {
			ll.Remove(current)
		}
		current = tempNext
	}
}

func (ll *DoublyLinkedList) Remove(node *Node) {
	if node == ll.Head {
		ll.Head = ll.Head.Next
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node == ll.Tail {
		ll.Tail = ll.Tail.Prev
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	node.Prev = nil
	node.Next = nil
}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	current := ll.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

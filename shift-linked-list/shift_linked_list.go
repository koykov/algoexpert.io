package shift_linked_list

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ShiftLinkedList(head *LinkedList, k int) *LinkedList {
	c, node_ := 1, head
	var tail *LinkedList
	for node_.Next != nil {
		if node_ = node_.Next; node_.Next == nil {
			tail = node_
		}
		c++
	}
	k = k % c
	if k < 0 {
		k = c + k
	}
	pnt := c - k

	node_ = head
	for i := 1; i < pnt; i++ {
		node_ = node_.Next
	}
	tail.Next = head
	head = node_.Next
	node_.Next = nil

	return head
}

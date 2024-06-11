package reverse_linked_list

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ReverseLinkedList(prev *LinkedList) *LinkedList {
	next := prev.Next
	prev.Next = nil
	for next != nil {
		curr := next
		next = next.Next
		curr.Next = prev
		prev = curr
	}
	return prev
}

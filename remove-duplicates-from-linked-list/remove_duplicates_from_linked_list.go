package remove_duplicates_from_linked_list

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(head *LinkedList) *LinkedList {
	p := head.Value
	pnode := head
	for {
		node := pnode.Next
		if node == nil {
			break
		}
		if node.Value == p {
			pnode.Next = node.Next
			continue
		}
		pnode = node
		p = node.Value
	}
	return head
}

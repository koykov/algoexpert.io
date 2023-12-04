package remove_kth_node_from_end

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	var n int
	head1 := head
	for head1 != nil {
		head1 = head1.Next
		n++
	}

	prev := head
	if n == k {
		*prev = *head.Next
		return
	}
	for i := 0; i < n-k; i++ {
		prev = head
		if head = head.Next; head == nil {
			return
		}
	}
	prev.Next = head.Next
}

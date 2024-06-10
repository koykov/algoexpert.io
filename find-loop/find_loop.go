package find_loop

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func FindLoop(head *LinkedList) *LinkedList {
	l, r := head.Next.Next, head.Next
	// floyd cycle
	for l != r {
		r = r.Next
		l = l.Next.Next
	}
	l = head
	for l != r {
		l = l.Next
		r = r.Next
	}
	return l
}

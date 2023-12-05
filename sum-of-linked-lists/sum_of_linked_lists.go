package sum_of_linked_lists

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func SumOfLinkedLists(a *LinkedList, b *LinkedList) *LinkedList {
	as := sum(a)
	bs := sum(b)
	s := as + bs
	sh := &LinkedList{}
	n := sh
	for {
		n.Value = s % 10
		s /= 10
		if s == 0 {
			break
		}
		n.Next = &LinkedList{}
		n = n.Next
	}
	return sh
}

func sum(head *LinkedList) (n int) {
	if head.Next != nil {
		n = sum(head.Next)
	}
	n = n*10 + head.Value
	return
}

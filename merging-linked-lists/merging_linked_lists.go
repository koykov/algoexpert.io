package merging_linked_lists

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MergingLinkedLists(a *LinkedList, b *LinkedList) *LinkedList {
	al, bl := lenOf(a), lenOf(b)

	switch {
	case al > bl:
		for i := 0; i < al-bl; i++ {
			a = a.Next
		}
	case bl > al:
		for i := 0; i < bl-al; i++ {
			b = b.Next
		}
	}

	if a == b && a != nil {
		return a
	}
	if a == nil || b == nil {
		return nil
	}
	return MergingLinkedLists(a.Next, b.Next)
}

func lenOf(l *LinkedList) int {
	if l == nil {
		return 0
	}
	return lenOf(l.Next) + 1
}

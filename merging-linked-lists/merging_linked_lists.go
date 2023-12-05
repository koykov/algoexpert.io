package merging_linked_lists

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MergingLinkedLists(a *LinkedList, b *LinkedList) *LinkedList {
	if a == b && a != nil {
		return a
	}
	if a == nil || b == nil {
		return nil
	}
	return MergingLinkedLists(a.Next, b.Next)
}

package merge_linked_lists

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MergeLinkedLists(h0 *LinkedList, h1 *LinkedList) *LinkedList {
	if h0.Value > h1.Value {
		h0, h1 = h1, h0
	}
	r := h0
	for {
		for h0 != nil && h1 != nil && h0.Value <= h1.Value {
			t := h0.Next
			h0.Next = h1
			h0.Next.Next = t
			h0, h1 = h0.Next, h1.Next
		}
		for h0 != nil && h1 != nil && h0.Value > h1.Value {
			t := h1.Next
			h1.Next = h0
			h1.Next.Next = t
			h0, h1 = h0.Next, h1.Next
		}
		break
	}
	return r
}

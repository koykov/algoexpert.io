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
		eq := h0.Value == h1.Value
		for h1 != nil && h0.Value == h1.Value {
			t := h1
			h1 = h1.Next
			t.Next = h0.Next
			h0.Next = t
		}
		for eq && h0.Value == h0.Next.Value {
			h0 = h0.Next
		}
		if h0.Value+1 == h1.Value {
			t := h1
			h1 = h1.Next
			t.Next = h0.Next
			h0.Next = t
		}
		h0 = h0.Next
		switch {
		case h0.Next == nil:
			h0.Next = h1
			goto exit
		case h1 == nil:
			goto exit
		}
	}
exit:
	_ = r
	return r
}

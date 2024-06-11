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
	c := h0
	if h0 = h0.Next; h0 == nil {
		c.Next = h1
		goto exit
	}
	for {
		if h0.Value <= h1.Value {
			c.Next = h0
			h0 = h0.Next
		} else {
			c.Next = h1
			h1 = h1.Next
		}
		c = c.Next
		switch {
		case h0 == nil:
			c.Next = h1
			goto exit
		case h1 == nil:
			c.Next = h0
			goto exit
		}
	}
exit:
	_ = r
	return r
}

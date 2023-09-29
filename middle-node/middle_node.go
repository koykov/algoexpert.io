package middle_node

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MiddleNode(head *LinkedList) *LinkedList {
	var c int
	node := head
	for node != nil {
		node = node.Next
		c++
	}
	if c == 1 {
		return head
	}
	node = head
	for i := 0; i < c/2; i++ {
		node = node.Next
	}
	return node
}

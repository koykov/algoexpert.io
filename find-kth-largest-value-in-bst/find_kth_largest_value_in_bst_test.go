package find_kth_largest_value_in_bst

import (
	"testing"
)

func TestBST(t *testing.T) {
	tree := freshBST()
	r := FindKthLargestValueInBst(tree, 3)
	if r != 17 {
		t.FailNow()
	}
}

func BenchmarkBST(b *testing.B) {
	b.ReportAllocs()
	tree := freshBST()
	for i := 0; i < b.N; i++ {
		FindKthLargestValueInBst(tree, 3)
	}
}

func newBST(value int) *BST {
	return &BST{Value: value}
}

func freshBST() *BST {
	root := newBST(10)
	root.Left = newBST(5)
	root.Left.Left = newBST(2)
	root.Left.Left.Left = newBST(1)
	root.Left.Left.Right = newBST(3)
	root.Left.Right = newBST(5)
	root.Right = newBST(20)
	root.Right.Left = newBST(17)
	root.Right.Right = newBST(22)

	return root
}

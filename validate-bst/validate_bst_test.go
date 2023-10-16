package validate_bst

import "testing"

func TestBST(t *testing.T) {
	tree := freshBST()
	if !tree.ValidateBst() {
		t.FailNow()
	}
}

func BenchmarkBST(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tree := freshBST()
		tree.ValidateBst()
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
	root.Left.Right = newBST(5)
	root.Right = newBST(15)
	root.Right.Left = newBST(13)
	root.Right.Left.Right = newBST(14)
	root.Right.Right = newBST(22)

	return root
}
